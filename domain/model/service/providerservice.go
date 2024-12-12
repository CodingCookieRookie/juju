// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"context"

	coremodel "github.com/juju/juju/core/model"
)

// ProviderState is the model state required by the provide service. This is
// the model database state, not the controller state.
type ProviderState interface {
	// Model returns a read-only model.
	GetModel(context.Context) (coremodel.ReadOnlyModel, error)
}

// ProviderService defines a service for interacting with the underlying model
// state, as opposed to the controller state.
type ProviderService struct {
	st ProviderState
}

// NewProviderService returns a new Service for interacting with a models state.
func NewProviderService(st ProviderState) *ProviderService {
	return &ProviderService{
		st: st,
	}
}

// Model returns a read-only model.
//
// The following error types can be expected to be returned:
// - [modelerrors.NotFound]: When the model is not found for a given uuid.
func (s *ProviderService) Model(ctx context.Context) (coremodel.ReadOnlyModel, error) {
	return s.st.GetModel(ctx)
}
