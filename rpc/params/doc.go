// Copyright 2025 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

// Package params contains the definition of the request and response parameters for the Juju API Server and
// Juju API callers (which includes the user-facing clients but also the agents and the controller itself).
// When introducing a change to a facade, to maintain compatibility with a previous version of the facade, copy the existing parameter struct
// and append to its name the version of the compatible facade (e.g., [params.AddApplicationUnitsV5](https://github.com/juju/juju/blob/e746c7c7fbf6d1e9be0807357c5bdcc210ec22d7/rpc/params/params.go#L451)).
// The new facade will use the undecorated name (e.g., [params.AddApplicationUnits](https://github.com/juju/juju/blob/e746c7c7fbf6d1e9be0807357c5bdcc210ec22d7/rpc/params/params.go#L441).

// API messages are encoded in JSON. The type `rpc.Request` represents
// a remote procedure call to be performed, absent its parameters. Those
// can be found in the type `rpc.Call`, which represents an active
// remote procedure call and embeds the request as well as other important
// parts of the request.

// Key components:

// - Defines constants for handling various API cases.
// - Defines the structs with json tags representing Juju entities to facilitate API interactions.

package params
