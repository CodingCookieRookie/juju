// Copyright 2025 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package removal

import (
	"context"

	"github.com/juju/worker/v4"
	"github.com/juju/worker/v4/dependency"

	coredependency "github.com/juju/juju/core/dependency"
	coreerrors "github.com/juju/juju/core/errors"
	"github.com/juju/juju/core/logger"
	"github.com/juju/juju/core/watcher"
	removalservice "github.com/juju/juju/domain/removal/service"
	"github.com/juju/juju/internal/errors"
	"github.com/juju/juju/internal/services"
)

// DomainServices describes the service factory
// required for working with entity removals.
type DomainServices interface {
	// Removal returns the service for managing entity removal.
	Removal() *removalservice.WatchableService
}

// RemovalService describes the ability to watch
// for and execute model entity removals.
type RemovalService interface {
	// WatchRemovals emits notifications for additions
	// and changes to removal job scheduling.
	WatchRemovals() (watcher.StringsWatcher, error)
}

// ManifoldConfig contains the configuration passed to this
// worker's manifold when run by the dependency engine.
type ManifoldConfig struct {

	// DomainServicesName is the name of the domain service factory dependency.
	DomainServicesName string

	// GetRemovalService is used to extract the removal
	// service from domain service dependency.
	GetRemovalService func(getter dependency.Getter, name string) (RemovalService, error)

	// NewWorker creates and returns a removal worker.
	NewWorker func(Config) (worker.Worker, error)

	// Logger logs stuff.
	Logger logger.Logger
}

// Validate ensures that the configuration is
// correctly populated for manifold operation.
func (config ManifoldConfig) Validate() error {
	if config.DomainServicesName == "" {
		return errors.New("empty DomainServicesName not valid").Add(coreerrors.NotValid)
	}
	if config.GetRemovalService == nil {
		return errors.New("nil GetRemovalService not valid").Add(coreerrors.NotValid)
	}
	if config.NewWorker == nil {
		return errors.New("nil NewWorker not valid").Add(coreerrors.NotValid)
	}
	if config.Logger == nil {
		return errors.New("nil Logger not valid").Add(coreerrors.NotValid)
	}

	return nil
}

// Manifold returns a dependency.Manifold that will run the removal worker.
func Manifold(config ManifoldConfig) dependency.Manifold {
	return dependency.Manifold{
		Inputs: []string{
			config.DomainServicesName,
		},
		Start: config.start,
	}
}

func (config ManifoldConfig) start(ctx context.Context, getter dependency.Getter) (worker.Worker, error) {
	if err := config.Validate(); err != nil {
		return nil, errors.Capture(err)
	}

	removalService, err := config.GetRemovalService(getter, config.DomainServicesName)
	if err != nil {
		return nil, errors.Capture(err)
	}

	wCfg := Config{
		RemovalService: removalService,
		Logger:         config.Logger,
	}

	w, err := config.NewWorker(wCfg)
	if err != nil {
		return nil, errors.Errorf("creating removal worker: %w", err)
	}
	return w, nil
}

// GetRemovalService extracts the model service factory from the input
// dependency getter, then returns the removal service from it.
func GetRemovalService(getter dependency.Getter, name string) (RemovalService, error) {
	return coredependency.GetDependencyByName(getter, name, func(factory services.ModelDomainServices) RemovalService {
		return factory.Removal()
	})
}
