// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package client

import (
	"context"

	"github.com/juju/juju/core/application"
	"github.com/juju/juju/core/arch"
	"github.com/juju/juju/core/blockdevice"
	"github.com/juju/juju/core/instance"
	"github.com/juju/juju/core/machine"
	"github.com/juju/juju/core/model"
	"github.com/juju/juju/core/network"
	"github.com/juju/juju/core/relation"
	"github.com/juju/juju/core/status"
	"github.com/juju/juju/core/unit"
	"github.com/juju/juju/domain/application/charm"
	domainmodel "github.com/juju/juju/domain/model"
	"github.com/juju/juju/domain/port"
	internalrelation "github.com/juju/juju/internal/relation"
)

// ApplicationService defines the methods that the facade assumes from the
// Application service.
type ApplicationService interface {
	// GetUnitUUID returns the UUID for the named unit
	GetUnitUUID(context.Context, unit.Name) (unit.UUID, error)

	// GetLatestPendingCharmhubCharm returns the latest charm that is pending
	// from the charmhub store. If there are no charms, returns is not found, as
	// [applicationerrors.CharmNotFound]. If there are multiple charms, then the
	// latest created at date is returned first.
	GetLatestPendingCharmhubCharm(ctx context.Context, name string, arch arch.Arch) (charm.CharmLocator, error)

	// GetApplicationIDByName returns an application ID by application name. It
	// returns an error if the application can not be found by the name. Returns
	// [applicationerrors.ApplicationNotFound] if the application is not found.
	GetApplicationIDByName(ctx context.Context, name string) (application.ID, error)

	// GetApplicationDisplayStatus returns the display status of the specified
	// application. The display status is equal to the application status if it
	// is set, otherwise it is derived from the unit display statuses. If no
	// application is found, an error satisfying
	// [applicationerrors.ApplicationNotFound] is returned.
	GetApplicationDisplayStatus(context.Context, application.ID) (*status.StatusInfo, error)

	// GetUnitDisplayStatus returns the display status of the specified unit.
	// The display status a function of both the unit workload status and the
	// cloud container status. It returns an error satisfying
	// [applicationerrors.UnitNotFound] if the unit doesn't exist.
	GetUnitDisplayStatus(context.Context, unit.Name) (*status.StatusInfo, error)

	// GetUnitAgentStatus returns the agent status of the specified unit. It
	// returns an error satisfying [applicationerrors.UnitNotFound] if the unit
	// doesn't exist.
	GetUnitAgentStatus(context.Context, unit.Name) (*status.StatusInfo, error)
}

// BlockDeviceService instances can fetch block devices for a machine.
type BlockDeviceService interface {
	BlockDevices(ctx context.Context, machineId string) ([]blockdevice.BlockDevice, error)
}

// MachineService defines the methods that the facade assumes from the Machine
// service.
type MachineService interface {
	// GetMachineUUID returns the UUID of a machine identified by its name.
	GetMachineUUID(ctx context.Context, name machine.Name) (string, error)
	// InstanceID returns the cloud specific instance id for this machine.
	InstanceID(ctx context.Context, machineUUID string) (instance.Id, error)
	// InstanceIDAndName returns the cloud specific instance ID and display name
	// for this machine.
	InstanceIDAndName(ctx context.Context, machineUUID string) (instance.Id, string, error)
	// HardwareCharacteristics returns the hardware characteristics of the
	// specified machine.
	HardwareCharacteristics(ctx context.Context, machineUUID string) (*instance.HardwareCharacteristics, error)
	// AppliedLXDProfiles returns the names of the LXD profiles on the machine.
	AppliedLXDProfileNames(ctx context.Context, machineUUID string) ([]string, error)
}

// ModelInfoService provides access to information about the model.
type ModelInfoService interface {
	// GetModelInfo returns information about the current model.
	GetModelInfo(context.Context) (model.ModelInfo, error)
	// GetStatus returns the current status of the model.
	// The following error types can be expected to be returned:
	// - [modelerrors.NotFound]: When the model does not exist.
	GetStatus(context.Context) (domainmodel.StatusInfo, error)
}

// NetworkService is the interface that is used to interact with the
// network spaces/subnets.
type NetworkService interface {
	// GetAllSpaces returns all spaces for the model.
	GetAllSpaces(ctx context.Context) (network.SpaceInfos, error)
	// GetAllSubnets returns all the subnets for the model.
	GetAllSubnets(ctx context.Context) (network.SubnetInfos, error)
}

// PortService defines the methods that the facade assumes from the Port
// service.
type PortService interface {
	// GetAllOpenedPorts returns the opened ports in the model, grouped by unit
	// name.
	GetAllOpenedPorts(ctx context.Context) (port.UnitGroupedPortRanges, error)

	// GetUnitOpenedPorts returns the opened ports for a given unit uuid,
	// grouped by endpoint.
	GetUnitOpenedPorts(context.Context, unit.UUID) (network.GroupedPortRanges, error)
}

// RelationService provides methods to interact with and retrieve details of
// relations within a model.
type RelationService interface {

	// AllRelations return all uuid of all relation for the current model.
	AllRelations(ctx context.Context) ([]relation.UUID, error)

	// GetRelationID returns the relation ID for the given relation UUID.
	GetRelationID(ctx context.Context, uuid relation.UUID) (int, error)

	// GetRelationEndpoints returns all endpoints for the given relation UUID
	GetRelationEndpoints(ctx context.Context, id relation.UUID) ([]internalrelation.Endpoint, error)

	// GetRelationStatus returns the status of the given relation.
	GetRelationStatus(ctx context.Context, relationUUID relation.UUID) (status.StatusInfo, error)
}
