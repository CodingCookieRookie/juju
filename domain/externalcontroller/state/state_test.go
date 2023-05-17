// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	ctx "context"

	"github.com/juju/collections/set"

	"github.com/juju/names/v4"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/utils/v3"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/core/crossmodel"
	"github.com/juju/juju/database/testing"
)

type stateSuite struct {
	testing.ControllerSuite
}

var _ = gc.Suite(&stateSuite{})

func (s *stateSuite) TestUpdateExternalControllerNewData(c *gc.C) {
	st := NewState(testing.TrackedDBFactory(s.TrackedDB()))

	ecUUID := utils.MustNewUUID().String()
	ec := crossmodel.ControllerInfo{
		ControllerTag: names.NewControllerTag(ecUUID),
		Alias:         "new-external-controller",
		Addrs:         []string{"10.10.10.10", "192.168.0.9"},
		CACert:        "random-cert-string",
	}

	m1 := utils.MustNewUUID().String()

	err := st.UpdateExternalController(ctx.Background(), ec, []string{m1})
	c.Assert(err, jc.ErrorIsNil)

	db := s.DB()

	// Check the controller record.
	row := db.QueryRow("SELECT alias, ca_cert FROM external_controller WHERE uuid = ?", ecUUID)
	c.Assert(row.Err(), jc.ErrorIsNil)

	var alias, cert string
	err = row.Scan(&alias, &cert)
	c.Assert(err, jc.ErrorIsNil)
	c.Check(alias, gc.Equals, "new-external-controller")
	c.Check(cert, gc.Equals, "random-cert-string")

	// Check the addresses.
	rows, err := db.Query("SELECT address FROM external_controller_address WHERE controller_uuid = ?", ecUUID)
	c.Assert(err, jc.ErrorIsNil)

	addrs := set.NewStrings()
	for rows.Next() {
		var addr string
		err := rows.Scan(&addr)
		c.Assert(err, jc.ErrorIsNil)
		addrs.Add(addr)
	}
	c.Check(addrs.Values(), gc.HasLen, 2)
	c.Check(addrs.Contains("10.10.10.10"), jc.IsTrue)
	c.Check(addrs.Contains("192.168.0.9"), jc.IsTrue)

	// Check for the model.
	row = db.QueryRow("SELECT controller_uuid FROM external_model WHERE uuid = ?", m1)
	c.Assert(row.Err(), jc.ErrorIsNil)

	var mc string
	err = row.Scan(&mc)
	c.Assert(err, jc.ErrorIsNil)
	c.Check(mc, gc.Equals, ecUUID)
}

func (s *stateSuite) TestUpdateExternalControllerUpsertAndReplace(c *gc.C) {
	st := NewState(testing.TrackedDBFactory(s.TrackedDB()))

	ecUUID := utils.MustNewUUID().String()
	ec := crossmodel.ControllerInfo{
		ControllerTag: names.NewControllerTag(ecUUID),
		Alias:         "new-external-controller",
		Addrs:         []string{"10.10.10.10", "192.168.0.9"},
		CACert:        "random-cert-string",
	}

	// Initial values.
	err := st.UpdateExternalController(ctx.Background(), ec, nil)
	c.Assert(err, jc.ErrorIsNil)

	// Now with different alias and addresses.
	ec.Alias = "updated-external-controller"
	ec.Addrs = []string{"10.10.10.10", "192.168.0.10"}

	err = st.UpdateExternalController(ctx.Background(), ec, nil)
	c.Assert(err, jc.ErrorIsNil)

	db := s.DB()

	// Check the controller record.
	row := db.QueryRow("SELECT alias FROM external_controller WHERE uuid = ?", ecUUID)
	c.Assert(row.Err(), jc.ErrorIsNil)

	var alias string
	err = row.Scan(&alias)
	c.Assert(err, jc.ErrorIsNil)
	c.Check(alias, gc.Equals, "updated-external-controller")

	// Addresses should have one preserved and one replaced.
	rows, err := db.Query("SELECT address FROM external_controller_address WHERE controller_uuid = ?", ecUUID)
	c.Assert(err, jc.ErrorIsNil)

	addrs := set.NewStrings()
	for rows.Next() {
		var addr string
		err := rows.Scan(&addr)
		c.Assert(err, jc.ErrorIsNil)
		addrs.Add(addr)
	}
	c.Check(addrs.Values(), gc.HasLen, 2)
	c.Check(addrs.Contains("10.10.10.10"), jc.IsTrue)
	c.Check(addrs.Contains("192.168.0.10"), jc.IsTrue)
}

func (s *stateSuite) TestUpdateExternalControllerUpdateModel(c *gc.C) {
	st := NewState(testing.TrackedDBFactory(s.TrackedDB()))

	// This is an existing controller with a model reference.
	ec := crossmodel.ControllerInfo{
		ControllerTag: names.NewControllerTag(utils.MustNewUUID().String()),
		Alias:         "existing-external-controller",
		CACert:        "random-cert-string",
	}

	m1 := utils.MustNewUUID().String()

	err := st.UpdateExternalController(ctx.Background(), ec, []string{m1})
	c.Assert(err, jc.ErrorIsNil)

	// Now upload a new controller with the same model
	ecUUID := utils.MustNewUUID().String()
	ec = crossmodel.ControllerInfo{
		ControllerTag: names.NewControllerTag(ecUUID),
		Alias:         "new-external-controller",
		CACert:        "another-random-cert-string",
	}

	err = st.UpdateExternalController(ctx.Background(), ec, []string{m1})
	c.Assert(err, jc.ErrorIsNil)

	// Check that the model is indicated as being on the new controller.
	row := s.DB().QueryRow("SELECT controller_uuid FROM external_model WHERE uuid = ?", m1)
	c.Assert(row.Err(), jc.ErrorIsNil)

	var mc string
	err = row.Scan(&mc)
	c.Assert(err, jc.ErrorIsNil)
	c.Check(mc, gc.Equals, ecUUID)
}