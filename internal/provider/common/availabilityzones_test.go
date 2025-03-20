// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package common_test

import (
	"context"
	"fmt"

	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/core/instance"
	"github.com/juju/juju/core/network"
	"github.com/juju/juju/environs"
	"github.com/juju/juju/environs/envcontext"
	"github.com/juju/juju/environs/instances"
	"github.com/juju/juju/internal/provider/common"
	coretesting "github.com/juju/juju/internal/testing"
)

type AvailabilityZoneSuite struct {
	coretesting.FakeJujuXDGDataHomeSuite
	env mockZonedEnviron

	callCtx envcontext.ProviderCallContext
}

var _ = gc.Suite(&AvailabilityZoneSuite{})

func (s *AvailabilityZoneSuite) SetUpSuite(c *gc.C) {
	s.FakeJujuXDGDataHomeSuite.SetUpSuite(c)

	s.callCtx = envcontext.WithoutCredentialInvalidator(context.Background())
	allInstances := make([]instances.Instance, 3)
	for i := range allInstances {
		allInstances[i] = &mockInstance{id: fmt.Sprintf("inst%d", i)}
	}
	s.env.allInstances = func(ctx context.Context) ([]instances.Instance, error) {
		return allInstances, nil
	}

	availabilityZones := make(network.AvailabilityZones, 3)
	for i := range availabilityZones {
		availabilityZones[i] = &mockAvailabilityZone{
			name:      fmt.Sprintf("az%d", i),
			available: i > 0,
		}
	}
	s.env.availabilityZones = func(ctx context.Context) (network.AvailabilityZones, error) {
		return availabilityZones, nil
	}
}

func (s *AvailabilityZoneSuite) TestAvailabilityZoneAllocationsAllRunningInstances(c *gc.C) {
	var called int
	s.PatchValue(&s.env.instanceAvailabilityZoneNames, func(ctx envcontext.ProviderCallContext, ids []instance.Id) (map[instance.Id]string, error) {
		c.Assert(ids, gc.DeepEquals, []instance.Id{"inst0", "inst1", "inst2"})
		called++
		return map[instance.Id]string{
			"inst0": "az0",
			"inst1": "az1",
			"inst2": "az2",
		}, nil
	})
	zoneInstances, err := common.AvailabilityZoneAllocations(&s.env, s.callCtx, nil)
	c.Assert(called, gc.Equals, 1)
	c.Assert(err, jc.ErrorIsNil)
	// az0 is unavailable, so az1 and az2 come out as equal best;
	// az1 comes first due to lexicographical ordering on the name.
	c.Assert(zoneInstances, gc.DeepEquals, []common.AvailabilityZoneInstances{{
		ZoneName:  "az1",
		Instances: []instance.Id{"inst1"},
	}, {
		ZoneName:  "az2",
		Instances: []instance.Id{"inst2"},
	}})
}

func (s *AvailabilityZoneSuite) TestAvailabilityZoneAllocationsAllRunningInstancesErrors(c *gc.C) {
	resultErr := fmt.Errorf("oh noes")
	s.PatchValue(&s.env.allInstances, func(context.Context) ([]instances.Instance, error) {
		return nil, resultErr
	})
	zoneInstances, err := common.AvailabilityZoneAllocations(&s.env, s.callCtx, nil)
	c.Assert(err, gc.Equals, resultErr)
	c.Assert(zoneInstances, gc.HasLen, 0)
}

func (s *AvailabilityZoneSuite) TestAvailabilityZoneAllocationsPartialInstances(c *gc.C) {
	var called int
	s.PatchValue(&s.env.instanceAvailabilityZoneNames, func(ctx envcontext.ProviderCallContext, ids []instance.Id) (map[instance.Id]string, error) {
		c.Assert(ids, gc.DeepEquals, []instance.Id{"nichts", "inst1", "null", "inst2"})
		called++
		return map[instance.Id]string{"inst1": "az1", "inst2": "az1"}, environs.ErrPartialInstances
	})
	zoneInstances, err := common.AvailabilityZoneAllocations(&s.env, s.callCtx, []instance.Id{"nichts", "inst1", "null", "inst2"})
	c.Assert(called, gc.Equals, 1)
	c.Assert(err, jc.ErrorIsNil)
	// az2 has fewer instances, so comes first.
	c.Assert(zoneInstances, gc.DeepEquals, []common.AvailabilityZoneInstances{{
		ZoneName: "az2",
	}, {
		ZoneName:  "az1",
		Instances: []instance.Id{"inst1", "inst2"},
	}})
}

func (s *AvailabilityZoneSuite) TestAvailabilityZoneAllocationsInstanceAvailabilityZonesErrors(c *gc.C) {
	returnErr := fmt.Errorf("whatever")
	var called int
	s.PatchValue(&s.env.instanceAvailabilityZoneNames, func(ctx envcontext.ProviderCallContext, ids []instance.Id) (map[instance.Id]string, error) {
		called++
		return nil, returnErr
	})
	zoneInstances, err := common.AvailabilityZoneAllocations(&s.env, s.callCtx, nil)
	c.Assert(called, gc.Equals, 1)
	c.Assert(err, gc.Equals, returnErr)
	c.Assert(zoneInstances, gc.HasLen, 0)
}

func (s *AvailabilityZoneSuite) TestAvailabilityZoneAllocationsInstanceAvailabilityZonesNoInstances(c *gc.C) {
	var called int
	s.PatchValue(&s.env.instanceAvailabilityZoneNames, func(ctx envcontext.ProviderCallContext, ids []instance.Id) (map[instance.Id]string, error) {
		called++
		return nil, environs.ErrNoInstances
	})
	zoneInstances, err := common.AvailabilityZoneAllocations(&s.env, s.callCtx, nil)
	c.Assert(called, gc.Equals, 1)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(zoneInstances, gc.HasLen, 2)
}

func (s *AvailabilityZoneSuite) TestAvailabilityZoneAllocationsNoZones(c *gc.C) {
	var calls []string
	s.PatchValue(&s.env.instanceAvailabilityZoneNames, func(ctx envcontext.ProviderCallContext, ids []instance.Id) (map[instance.Id]string, error) {
		c.Assert(ids, gc.DeepEquals, []instance.Id{"inst0", "inst1", "inst2"})
		calls = append(calls, "InstanceAvailabilityZoneNames")
		return make(map[instance.Id]string, 3), nil
	})
	s.PatchValue(&s.env.availabilityZones, func(context.Context) (network.AvailabilityZones, error) {
		calls = append(calls, "AvailabilityZones")
		return network.AvailabilityZones{}, nil
	})
	zoneInstances, err := common.AvailabilityZoneAllocations(&s.env, s.callCtx, nil)
	c.Assert(calls, gc.DeepEquals, []string{"InstanceAvailabilityZoneNames", "AvailabilityZones"})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(zoneInstances, gc.HasLen, 0)
}

func (s *AvailabilityZoneSuite) TestAvailabilityZoneAllocationsErrors(c *gc.C) {
	var calls []string
	s.PatchValue(&s.env.instanceAvailabilityZoneNames, func(ctx envcontext.ProviderCallContext, ids []instance.Id) (map[instance.Id]string, error) {
		c.Assert(ids, gc.DeepEquals, []instance.Id{"inst0", "inst1", "inst2"})
		calls = append(calls, "InstanceAvailabilityZoneNames")
		return make(map[instance.Id]string, 3), nil
	})
	resultErr := fmt.Errorf("u can haz no az")
	s.PatchValue(&s.env.availabilityZones, func(context.Context) (network.AvailabilityZones, error) {
		calls = append(calls, "AvailabilityZones")
		return nil, resultErr
	})
	zoneInstances, err := common.AvailabilityZoneAllocations(&s.env, s.callCtx, nil)
	c.Assert(calls, gc.DeepEquals, []string{"InstanceAvailabilityZoneNames", "AvailabilityZones"})
	c.Assert(err, gc.Equals, resultErr)
	c.Assert(zoneInstances, gc.HasLen, 0)
}
