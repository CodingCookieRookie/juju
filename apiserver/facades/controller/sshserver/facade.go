// Copyright 2025 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package sshserver

import (
	apiservererrors "github.com/juju/juju/apiserver/errors"
	"github.com/juju/juju/apiserver/facade"
	"github.com/juju/juju/controller"
	"github.com/juju/juju/rpc/params"
	"github.com/juju/juju/state"
	"github.com/juju/juju/state/watcher"
)

// Backend provides required state for the Facade.
type Backend interface {
	ControllerConfig() (controller.Config, error)
	WatchControllerConfig() state.NotifyWatcher
	SSHServerHostKey() (string, error)
}

// Facade allows model config manager clients to watch controller config changes and fetch controller config.
type Facade struct {
	resources facade.Resources

	backend Backend
}

// NewFacade returns a new SSHServer facade to be registered for use within
// the worker.
func NewFacade(ctx facade.Context, backend Backend) *Facade {
	return &Facade{
		resources: ctx.Resources(),
		backend:   backend,
	}
}

// ControllerConfig returns the current controller config.
func (f *Facade) ControllerConfig() (params.ControllerConfigResult, error) {
	result := params.ControllerConfigResult{}
	config, err := f.backend.ControllerConfig()
	if err != nil {
		return result, err
	}
	result.Config = params.ControllerConfig(config)
	return result, nil
}

// WatchControllerConfig creates a watcher and returns it's ID for watching upon.
func (f *Facade) WatchControllerConfig() (params.NotifyWatchResult, error) {
	result := params.NotifyWatchResult{}
	w := f.backend.WatchControllerConfig()
	if _, ok := <-w.Changes(); ok {
		result.NotifyWatcherId = f.resources.Register(w)
	} else {
		return result, watcher.EnsureErr(w)
	}
	return result, nil
}

// SSHServerHostKey returns the controller's SSH server host key.
func (f *Facade) SSHServerHostKey() (params.StringResult, error) {
	result := params.StringResult{}
	key, err := f.backend.SSHServerHostKey()
	if err != nil {
		result.Error = apiservererrors.ServerError(err)
	}
	result.Result = key
	return result, nil
}
