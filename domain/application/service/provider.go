// Copyright 2025 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"context"
	"strings"

	"github.com/juju/clock"

	coreapplication "github.com/juju/juju/core/application"
	"github.com/juju/juju/core/assumes"
	corecharm "github.com/juju/juju/core/charm"
	coreconstraints "github.com/juju/juju/core/constraints"
	coreerrors "github.com/juju/juju/core/errors"
	"github.com/juju/juju/core/leadership"
	"github.com/juju/juju/core/logger"
	coremodel "github.com/juju/juju/core/model"
	"github.com/juju/juju/core/providertracker"
	corestorage "github.com/juju/juju/core/storage"
	"github.com/juju/juju/domain/application"
	applicationerrors "github.com/juju/juju/domain/application/errors"
	"github.com/juju/juju/domain/constraints"
	modelerrors "github.com/juju/juju/domain/model/errors"
	storageerrors "github.com/juju/juju/domain/storage/errors"
	"github.com/juju/juju/environs/envcontext"
	internalcharm "github.com/juju/juju/internal/charm"
	"github.com/juju/juju/internal/errors"
	"github.com/juju/juju/internal/storage"
)

// ProviderService defines a service for interacting with the underlying
// model state.
type ProviderService struct {
	*Service

	modelID            coremodel.UUID
	agentVersionGetter AgentVersionGetter
	provider           providertracker.ProviderGetter[Provider]
	// This provider is separated from [provider] because the
	// [SupportedFeatureProvider] interface is only satisfied by the
	// k8s provider.
	supportedFeatureProvider providertracker.ProviderGetter[SupportedFeatureProvider]
}

// NewProviderService returns a new Service for interacting with a models state.
func NewProviderService(
	st State,
	leaderEnsurer leadership.Ensurer,
	storageRegistryGetter corestorage.ModelStorageRegistryGetter,
	modelID coremodel.UUID,
	agentVersionGetter AgentVersionGetter,
	provider providertracker.ProviderGetter[Provider],
	supportedFeatureProvider providertracker.ProviderGetter[SupportedFeatureProvider],
	charmStore CharmStore,
	statusHistory StatusHistory,
	clock clock.Clock,
	logger logger.Logger,
) *ProviderService {
	return &ProviderService{
		Service: NewService(
			st,
			leaderEnsurer,
			storageRegistryGetter,
			charmStore,
			statusHistory,
			clock,
			logger,
		),
		modelID:                  modelID,
		agentVersionGetter:       agentVersionGetter,
		provider:                 provider,
		supportedFeatureProvider: supportedFeatureProvider,
	}
}

func (s *Service) poolStorageProvider(
	ctx context.Context,
	registry storage.ProviderRegistry,
	poolNameOrType string,
) (storage.Provider, error) {
	pool, err := s.st.GetStoragePoolByName(ctx, poolNameOrType)
	if errors.Is(err, storageerrors.PoolNotFoundError) {
		// If there's no pool called poolNameOrType, maybe a provider type
		// has been specified directly.
		providerType := storage.ProviderType(poolNameOrType)
		aProvider, registryErr := registry.StorageProvider(providerType)
		if registryErr != nil {
			// The name can't be resolved as a storage provider type,
			// so return the original "pool not found" error.
			return nil, errors.Capture(err)
		}
		return aProvider, nil
	} else if err != nil {
		return nil, errors.Capture(err)
	}
	providerType := storage.ProviderType(pool.Provider)
	aProvider, err := registry.StorageProvider(providerType)
	if err != nil {
		return nil, errors.Capture(err)
	}
	return aProvider, nil
}

// CreateApplication creates the specified application and units if required,
// returning an error satisfying [applicationerrors.ApplicationAlreadyExists]
// if the application already exists.
func (s *ProviderService) CreateApplication(
	ctx context.Context,
	name string,
	charm internalcharm.Charm,
	origin corecharm.Origin,
	args AddApplicationArgs,
	units ...AddUnitArg,
) (coreapplication.ID, error) {
	if err := validateCharmAndApplicationParams(
		name,
		args.ReferenceName,
		charm,
		origin,
		args.DownloadInfo,
	); err != nil {
		return "", errors.Errorf("invalid application args: %w", err)
	}

	if err := validateCreateApplicationResourceParams(charm, args.ResolvedResources, args.PendingResources); err != nil {
		return "", errors.Errorf("create application: %w", err)
	}

	modelType, err := s.st.GetModelType(ctx)
	if err != nil {
		return "", errors.Errorf("getting model type: %w", err)
	}
	appArg, err := makeCreateApplicationArgs(ctx, s.st, s.storageRegistryGetter, modelType, charm, origin, args)
	if err != nil {
		return "", errors.Errorf("creating application args: %w", err)
	}
	// We know that the charm name is valid, so we can use it as the application
	// name if that is not provided.
	if name == "" {
		// Annoyingly this should be the reference name, but that's not
		// true in the previous code. To keep compatibility, we'll use the
		// charm name.
		name = appArg.Charm.Metadata.Name
	}

	numUnits := len(units)
	appArg.Scale = numUnits

	cons, err := s.mergeApplicationAndModelConstraints(ctx, constraints.DecodeConstraints(args.Constraints))
	if err != nil {
		return "", errors.Errorf("merging application and model constraints: %w", err)
	}

	// Adding units with storage needs to know the kind of storage supported
	// by the underlying provider so gather that here as it needs to be
	// done outside a transaction.
	registry, err := s.storageRegistryGetter.GetStorageRegistry(ctx)
	if err != nil {
		return "", err
	}

	unitArgs, err := s.makeUnitArgs(modelType, units, constraints.DecodeConstraints(cons))
	if err != nil {
		return "", errors.Errorf("making unit args: %w", err)
	}

	if len(appArg.Storage) > 0 {
		appArg.StoragePoolKind = make(map[string]storage.StorageKind)
	}
	for _, arg := range appArg.Storage {
		p, err := s.poolStorageProvider(ctx, registry, arg.PoolNameOrType)
		if err != nil {
			return "", err
		}
		if p.Supports(storage.StorageKindFilesystem) {
			appArg.StoragePoolKind[arg.PoolNameOrType] = storage.StorageKindFilesystem
		}
		if p.Supports(storage.StorageKindBlock) {
			appArg.StoragePoolKind[arg.PoolNameOrType] = storage.StorageKindBlock
		}
	}
	appID, err := s.st.CreateApplication(ctx, name, appArg, unitArgs)
	if err != nil {
		return "", errors.Errorf("creating application %q: %w", name, err)
	}

	s.logger.Infof(ctx, "created application %q with ID %q", name, appID)

	if args.ApplicationStatus != nil {
		if err := s.statusHistory.RecordStatus(ctx, applicationNamespace.WithID(appID.String()), *args.ApplicationStatus); err != nil {
			s.logger.Infof(ctx, "failed recording application status history: %w", err)
		}
	}

	return appID, nil
}

// GetSupportedFeatures returns the set of features that the model makes
// available for charms to use.
// If the agent version cannot be found, an error satisfying
// [modelerrors.NotFound] will be returned.
func (s *ProviderService) GetSupportedFeatures(ctx context.Context) (assumes.FeatureSet, error) {
	agentVersion, err := s.agentVersionGetter.GetTargetAgentVersion(ctx)
	if err != nil {
		return assumes.FeatureSet{}, err
	}

	var fs assumes.FeatureSet
	fs.Add(assumes.Feature{
		Name:        "juju",
		Description: assumes.UserFriendlyFeatureDescriptions["juju"],
		Version:     &agentVersion,
	})

	supportedFeatureProvider, err := s.supportedFeatureProvider(ctx)
	if errors.Is(err, coreerrors.NotSupported) {
		return fs, nil
	} else if err != nil {
		return fs, err
	}

	envFs, err := supportedFeatureProvider.SupportedFeatures()
	if err != nil {
		return fs, errors.Errorf("enumerating features supported by environment: %w", err)
	}

	fs.Merge(envFs)

	return fs, nil
}

// SetApplicationConstraints sets the application constraints for the
// specified application ID.
// This method overwrites the full constraints on every call.
// If invalid constraints are provided (e.g. invalid container type or
// non-existing space), a [applicationerrors.InvalidApplicationConstraints]
// error is returned.
// If no application is found, an error satisfying
// [applicationerrors.ApplicationNotFound] is returned.
func (s *ProviderService) SetApplicationConstraints(ctx context.Context, appID coreapplication.ID, cons coreconstraints.Value) error {
	if err := appID.Validate(); err != nil {
		return errors.Errorf("application ID: %w", err)
	}
	if err := s.validateConstraints(ctx, cons); err != nil {
		return err
	}

	return s.st.SetApplicationConstraints(ctx, appID, constraints.DecodeConstraints(cons))
}

func (s *ProviderService) constraintsValidator(ctx context.Context) (coreconstraints.Validator, error) {
	provider, err := s.provider(ctx)
	if errors.Is(err, coreerrors.NotSupported) {
		// Not validating constraints, as the provider doesn't support it.
		return nil, nil
	} else if err != nil {
		return nil, errors.Capture(err)
	}

	validator, err := provider.ConstraintsValidator(envcontext.WithoutCredentialInvalidator(ctx))
	if err != nil {
		return nil, errors.Capture(err)
	}

	return validator, nil
}

// AddUnits adds the specified units to the application, returning an error
// satisfying [applicationerrors.ApplicationNotFoundError] if the application
// doesn't exist.
// If no units are provided, it will return nil.
func (s *ProviderService) AddUnits(ctx context.Context, storageParentDir, appName string, units ...AddUnitArg) error {
	if !isValidApplicationName(appName) {
		return applicationerrors.ApplicationNameNotValid
	}

	if len(units) == 0 {
		return nil
	}

	appUUID, err := s.st.GetApplicationIDByName(ctx, appName)
	if err != nil {
		return errors.Errorf("getting application %q id: %w", appName, err)
	}

	modelType, err := s.st.GetModelType(ctx)
	if err != nil {
		return errors.Errorf("getting model type: %w", err)
	}

	appCons, err := s.st.GetApplicationConstraints(ctx, appUUID)
	if err != nil {
		return errors.Errorf("getting application %q constraints: %w", appName, err)
	}

	cons, err := s.mergeApplicationAndModelConstraints(ctx, appCons)
	if err != nil {
		return errors.Capture(err)
	}

	args, err := s.makeUnitArgs(modelType, units, constraints.DecodeConstraints(cons))
	if err != nil {
		return errors.Errorf("making unit args: %w", err)
	}

	if modelType == coremodel.IAAS {
		err = s.st.AddIAASUnits(ctx, storageParentDir, appUUID, args...)
	} else {
		err = s.st.AddCAASUnits(ctx, storageParentDir, appUUID, args...)
	}
	if err != nil {
		return errors.Errorf("adding units to application %q: %w", appName, err)
	}

	for _, arg := range args {
		unitName := arg.UnitName.String()

		// The agent and workload status are required to be provided when adding
		// a unit.
		if arg.AgentStatus == nil || arg.WorkloadStatus == nil {
			return errors.Errorf("unit %q status not provided", unitName)
		}

		// Force the presence to be recorded as true, as the unit has just been
		// added.
		if agentStatus, err := decodeUnitAgentStatus(&application.UnitStatusInfo[application.UnitAgentStatusType]{
			StatusInfo: *arg.AgentStatus,
			Present:    true,
		}); err == nil && agentStatus != nil {
			if err := s.statusHistory.RecordStatus(ctx, unitAgentNamespace.WithID(unitName), *agentStatus); err != nil {
				s.logger.Infof(ctx, "failed recording agent status for unit %q: %v", unitName, err)
			}
		}

		if workloadStatus, err := decodeUnitWorkloadStatus(&application.UnitStatusInfo[application.WorkloadStatusType]{
			StatusInfo: *arg.WorkloadStatus,
			Present:    true,
		}); err == nil && workloadStatus != nil {
			if err := s.statusHistory.RecordStatus(ctx, unitWorkloadNamespace.WithID(unitName), *workloadStatus); err != nil {
				s.logger.Infof(ctx, "failed recording workload status for unit %q: %v", unitName, err)
			}
		}
	}

	return nil
}

func (s *ProviderService) mergeApplicationAndModelConstraints(ctx context.Context, appCons constraints.Constraints) (coreconstraints.Value, error) {
	// If the provider doesn't support constraints validation, then we can
	// just return the zero value.
	validator, err := s.constraintsValidator(ctx)
	if err != nil || validator == nil {
		return coreconstraints.Value{}, errors.Capture(err)
	}

	modelCons, err := s.st.GetModelConstraints(ctx)
	if err != nil && !errors.Is(err, modelerrors.ConstraintsNotFound) {
		return coreconstraints.Value{}, errors.Errorf("retrieving model constraints constraints: %w	", err)
	}

	res, err := validator.Merge(constraints.EncodeConstraints(appCons), constraints.EncodeConstraints(modelCons))
	if err != nil {
		return coreconstraints.Value{}, errors.Errorf("merging application and model constraints: %w", err)
	}

	return res, nil
}

func (s *ProviderService) validateConstraints(ctx context.Context, cons coreconstraints.Value) error {
	validator, err := s.constraintsValidator(ctx)
	if err != nil {
		return errors.Capture(err)
	} else if validator == nil {
		return nil
	}

	unsupported, err := validator.Validate(cons)
	if len(unsupported) > 0 {
		s.logger.Warningf(ctx,
			"unsupported constraints: %v", strings.Join(unsupported, ","))
	} else if err != nil {
		return errors.Capture(err)
	}

	return nil
}
