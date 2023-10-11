// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"context"
	"regexp"
	"time"

	"github.com/canonical/sqlair"
	"github.com/juju/errors"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/worker/v3/workertest"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/cloud"
	"github.com/juju/juju/core/changestream"
	"github.com/juju/juju/core/watcher"
	"github.com/juju/juju/core/watcher/eventsource"
	"github.com/juju/juju/core/watcher/watchertest"
	dbcloud "github.com/juju/juju/domain/cloud/state"
	changestreamtesting "github.com/juju/juju/internal/changestream/testing"
	jujutesting "github.com/juju/juju/testing"
)

type credentialSuite struct {
	changestreamtesting.ControllerSuite
}

var _ = gc.Suite(&credentialSuite{})

func (s *credentialSuite) SetUpTest(c *gc.C) {
	s.ControllerSuite.SetUpTest(c)

	s.addCloud(c, cloud.Cloud{
		Name:      "stratus",
		Type:      "ec2",
		AuthTypes: cloud.AuthTypes{cloud.AccessKeyAuthType, cloud.UserPassAuthType},
	})
}

func (s *credentialSuite) addCloud(c *gc.C, cloud cloud.Cloud) string {
	cloudSt := dbcloud.NewState(s.TxnRunnerFactory())
	ctx := context.Background()
	err := cloudSt.UpsertCloud(ctx, cloud)
	c.Assert(err, jc.ErrorIsNil)

	db := s.DB()
	row := db.QueryRow("SELECT uuid FROM cloud WHERE name = ?", cloud.Name)
	c.Assert(row.Err(), jc.ErrorIsNil)

	var dbCloud dbcloud.Cloud
	err = row.Scan(&dbCloud.ID)
	c.Assert(err, jc.ErrorIsNil)
	return dbCloud.ID
}

func (s *credentialSuite) TestUpdateCloudCredentialNew(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	cred := cloud.NewNamedCredential("foobar", cloud.AccessKeyAuthType, map[string]string{
		"foo": "foo val",
		"bar": "bar val",
	}, true)
	ctx := context.Background()
	err := st.UpsertCloudCredential(ctx, "foobar", "stratus", "bob", cred)
	c.Assert(err, jc.ErrorIsNil)

	out, err := st.CloudCredential(ctx, "foobar", "stratus", "bob")
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(out, jc.DeepEquals, cred)
}

func (s *credentialSuite) TestUpdateCloudCredentialNoValues(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	cred := cloud.NewNamedCredential("foobar", cloud.AccessKeyAuthType, map[string]string{}, true)
	ctx := context.Background()
	err := st.UpsertCloudCredential(ctx, "foobar", "stratus", "bob", cred)
	c.Assert(err, jc.ErrorIsNil)

	out, err := st.CloudCredential(ctx, "foobar", "stratus", "bob")
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(out, jc.DeepEquals, cred)
}

func (s *credentialSuite) TestUpdateCloudCredentialMissingName(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	cred := cloud.NewCredential(cloud.AccessKeyAuthType, map[string]string{
		"foo": "foo val",
		"bar": "bar val",
	})
	ctx := context.Background()
	err := st.UpsertCloudCredential(ctx, "", "stratus", "bob", cred)
	c.Assert(err, gc.ErrorMatches, "updating credential: credential name cannot be empty")
}

func (s *credentialSuite) TestCreateInvalidCredential(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	cred := cloud.NewCredential(cloud.AccessKeyAuthType, map[string]string{
		"foo": "foo val",
		"bar": "bar val",
	})
	// Setting of these properties should have no effect when creating a new credential.
	cred.Invalid = true
	cred.InvalidReason = "because am testing you"
	ctx := context.Background()
	err := st.UpsertCloudCredential(ctx, "foobar", "stratus", "bob", cred)
	c.Assert(err, gc.ErrorMatches, "adding invalid credential not supported")
}

func (s *credentialSuite) TestUpdateCloudCredentialExisting(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	cred := cloud.NewNamedCredential("foobar", cloud.AccessKeyAuthType, map[string]string{
		"foo": "foo val",
		"bar": "bar val",
	}, false)
	ctx := context.Background()
	err := st.UpsertCloudCredential(ctx, "foobar", "stratus", "bob", cred)
	c.Assert(err, jc.ErrorIsNil)

	cred = cloud.NewNamedCredential("foobar", cloud.UserPassAuthType, map[string]string{
		"user":     "bob's nephew",
		"password": "simple",
	}, true)
	err = st.UpsertCloudCredential(ctx, "foobar", "stratus", "bob", cred)
	c.Assert(err, jc.ErrorIsNil)

	out, err := st.CloudCredential(ctx, "foobar", "stratus", "bob")
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(out, jc.DeepEquals, cred)
}

func (s *credentialSuite) TestUpdateCloudCredentialInvalidAuthType(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	cred := cloud.NewNamedCredential("foobar", cloud.OAuth2AuthType, map[string]string{
		"foo": "foo val",
		"bar": "bar val",
	}, false)
	ctx := context.Background()
	err := st.UpsertCloudCredential(ctx, "foobar", "stratus", "bob", cred)
	c.Assert(err, gc.ErrorMatches, regexp.QuoteMeta(
		`updating credential: validating credential "foobar" owned by "bob" for cloud "stratus": supported auth-types ["access-key" "userpass"], "oauth2" not supported`))
}

func (s *credentialSuite) TestCloudCredentialsEmpty(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	creds, err := st.CloudCredentials(context.Background(), "bob", "dummy")
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(creds, gc.HasLen, 0)
}

func (s *credentialSuite) TestCloudCredentials(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	cred1 := cloud.NewCredential(cloud.AccessKeyAuthType, map[string]string{
		"foo": "foo val",
		"bar": "bar val",
	})
	ctx := context.Background()
	err := st.UpsertCloudCredential(ctx, "bobcred1", "stratus", "bob", cred1)
	c.Assert(err, jc.ErrorIsNil)

	cred2 := cloud.NewCredential(cloud.AccessKeyAuthType, map[string]string{
		"baz": "baz val",
		"qux": "qux val",
	})
	err = st.UpsertCloudCredential(ctx, "bobcred2", "stratus", "bob", cred2)
	c.Assert(err, jc.ErrorIsNil)
	err = st.UpsertCloudCredential(ctx, "foobar", "stratus", "mary", cred2)
	c.Assert(err, jc.ErrorIsNil)

	cred1.Label = "bobcred1"
	cred2.Label = "bobcred2"

	for _, userName := range []string{"bob", "bob"} {
		creds, err := st.CloudCredentials(ctx, userName, "stratus")
		c.Assert(err, jc.ErrorIsNil)
		c.Assert(creds, jc.DeepEquals, map[string]cloud.Credential{
			"stratus/bob/bobcred1": cred1,
			"stratus/bob/bobcred2": cred2,
		})
	}
}

func (s *credentialSuite) assertCredentialInvalidated(c *gc.C, st *State, cloudName, userName, credentialName string) {
	cred := cloud.NewCredential(cloud.AccessKeyAuthType, map[string]string{
		"foo": "foo val",
		"bar": "bar val",
	})
	ctx := context.Background()
	err := st.UpsertCloudCredential(ctx, credentialName, cloudName, userName, cred)
	c.Assert(err, jc.ErrorIsNil)

	cred = cloud.NewCredential(cloud.UserPassAuthType, map[string]string{
		"user":     "bob's nephew",
		"password": "simple",
	})
	cred.Invalid = true
	cred.InvalidReason = "because it is really really invalid"
	err = st.UpsertCloudCredential(ctx, credentialName, cloudName, userName, cred)
	c.Assert(err, jc.ErrorIsNil)

	out, err := st.CloudCredential(ctx, credentialName, cloudName, userName)
	c.Assert(err, jc.ErrorIsNil)
	cred.Label = "foobar"
	c.Assert(out, jc.DeepEquals, cred)
}

func (s *credentialSuite) TestInvalidateCredential(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())
	s.assertCredentialInvalidated(c, st, "stratus", "bob", "foobar")
}

func (s *credentialSuite) assertCredentialMarkedValid(c *gc.C, st *State, cloudName, userName, credentialName string, credential cloud.Credential) {
	ctx := context.Background()
	err := st.UpsertCloudCredential(ctx, credentialName, cloudName, userName, credential)
	c.Assert(err, jc.ErrorIsNil)

	out, err := st.CloudCredential(ctx, credentialName, cloudName, userName)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(out.Invalid, jc.IsFalse)
}

func (s *credentialSuite) TestMarkInvalidCredentialAsValidExplicitly(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())
	// This call will ensure that there is an invalid credential to test with.
	s.assertCredentialInvalidated(c, st, "stratus", "bob", "foobar")

	cred := cloud.NewCredential(cloud.UserPassAuthType, map[string]string{
		"user":     "bob's nephew",
		"password": "simple",
	})
	cred.Invalid = false
	s.assertCredentialMarkedValid(c, st, "stratus", "bob", "foobar", cred)
}

func (s *credentialSuite) TestMarkInvalidCredentialAsValidImplicitly(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())
	// This call will ensure that there is an invalid credential to test with.
	s.assertCredentialInvalidated(c, st, "stratus", "bob", "foobar")

	cred := cloud.NewCredential(cloud.UserPassAuthType, map[string]string{
		"user":     "bob's nephew",
		"password": "simple",
	})
	s.assertCredentialMarkedValid(c, st, "stratus", "bob", "foobar", cred)
}

func (s *credentialSuite) TestRemoveCredentials(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	cred1 := cloud.NewCredential(cloud.AccessKeyAuthType, map[string]string{
		"foo": "foo val",
		"bar": "bar val",
	})
	ctx := context.Background()
	err := st.UpsertCloudCredential(ctx, "bobcred1", "stratus", "bob", cred1)
	c.Assert(err, jc.ErrorIsNil)

	err = st.RemoveCloudCredential(ctx, "bobcred1", "stratus", "bob")
	c.Assert(err, jc.ErrorIsNil)

	_, err = st.CloudCredential(ctx, "bobcred1", "stratus", "bob")
	c.Assert(err, jc.ErrorIs, errors.NotFound)
}

func (s *credentialSuite) TestAllCloudCredentialsNotFound(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	out, err := st.AllCloudCredentials(context.Background(), "bob")
	c.Assert(err, gc.ErrorMatches, "cloud credentials for \"bob\" not found")
	c.Assert(out, gc.IsNil)
}

func (s *credentialSuite) createCloudCredential(c *gc.C, st *State, credentialName, cloudName, userName string) cloud.Credential {
	authType := cloud.AccessKeyAuthType
	attributes := map[string]string{
		"foo": "foo val",
		"bar": "bar val",
	}

	s.addCloud(c, cloud.Cloud{
		Name:      cloudName,
		Type:      "ec2",
		AuthTypes: cloud.AuthTypes{cloud.AccessKeyAuthType, cloud.UserPassAuthType},
	})

	cred := cloud.NewNamedCredential(credentialName, authType, attributes, false)
	err := st.UpsertCloudCredential(context.Background(), credentialName, cloudName, userName, cred)
	c.Assert(err, jc.ErrorIsNil)
	return cred
}

func (s *credentialSuite) TestAllCloudCredentials(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	one := s.createCloudCredential(c, st, "foobar", "cirrus", "bob")
	two := s.createCloudCredential(c, st, "foobar", "stratus", "bob")

	// Added to make sure it is not returned.
	s.createCloudCredential(c, st, "foobar", "cumulus", "mary")

	out, err := st.AllCloudCredentials(context.Background(), "bob")
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(out, jc.DeepEquals, []CloudCredential{
		{Credential: one, CloudName: "cirrus"},
		{Credential: two, CloudName: "stratus"},
	})
}

func (s *credentialSuite) TestInvalidateCloudCredential(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	one := s.createCloudCredential(c, st, "foobar", "cirrus", "bob")
	c.Assert(one.Invalid, jc.IsFalse)

	ctx := context.Background()
	reason := "testing, testing 1,2,3"
	err := st.InvalidateCloudCredential(ctx, "foobar", "cirrus", "bob", reason)
	c.Assert(err, jc.ErrorIsNil)

	updated, err := st.CloudCredential(ctx, "foobar", "cirrus", "bob")
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(updated.Invalid, jc.IsTrue)
	c.Assert(updated.InvalidReason, gc.Equals, reason)
}

func (s *credentialSuite) TestInvalidateCloudCredentialNotFound(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	ctx := context.Background()
	err := st.InvalidateCloudCredential(ctx, "foobar", "cirrus", "bob", "reason")
	c.Assert(err, jc.ErrorIs, errors.NotFound)
}

type watcherFunc func(namespace, changeValue string, changeMask changestream.ChangeType) (watcher.NotifyWatcher, error)

func (f watcherFunc) NewValueWatcher(
	namespace, changeValue string, changeMask changestream.ChangeType,
) (watcher.NotifyWatcher, error) {
	return f(namespace, changeValue, changeMask)
}

func (s *credentialSuite) watcherFunc(c *gc.C, expectedChangeValue string) watcherFunc {
	return func(namespace, changeValue string, changeMask changestream.ChangeType) (watcher.NotifyWatcher, error) {
		c.Assert(namespace, gc.Equals, "cloud_credential")
		c.Assert(changeMask, gc.Equals, changestream.All)
		c.Assert(changeValue, gc.Equals, expectedChangeValue)

		db, err := s.GetWatchableDB(namespace)
		c.Assert(err, jc.ErrorIsNil)

		base := eventsource.NewBaseWatcher(db, jujutesting.NewCheckLogger(c))
		return eventsource.NewValueWatcher(base, namespace, changeValue, changeMask), nil
	}
}

func (s *credentialSuite) TestWatchCredentialNotFound(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	ctx := context.Background()
	_, err := st.WatchCredential(ctx, s.watcherFunc(c, ""), "foobar", "cirrus", "bob")
	c.Assert(err, jc.ErrorIs, errors.NotFound)
}

func (s *credentialSuite) TestWatchCredential(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())
	s.createCloudCredential(c, st, "foobar", "cirrus", "bob")

	var uuid string
	err := s.TxnRunner().Txn(context.Background(), func(ctx context.Context, tx *sqlair.TX) error {
		var err error
		uuid, err = st.credentialUUID(ctx, tx, "foobar", "cirrus", "bob")
		return err
	})
	c.Assert(err, jc.ErrorIsNil)

	w, err := st.WatchCredential(context.Background(), s.watcherFunc(c, uuid), "foobar", "cirrus", "bob")
	c.Assert(err, jc.ErrorIsNil)
	s.AddCleanup(func(c *gc.C) { workertest.CleanKill(c, w) })

	wc := watchertest.NewNotifyWatcherC(c, w)
	wc.AssertChanges(time.Second) // Initial event.

	cred := cloud.NewNamedCredential("foobar", cloud.AccessKeyAuthType, map[string]string{
		"foo": "foo val",
		"bar": "bar val",
	}, true)
	err = s.TxnRunner().Txn(context.Background(), func(ctx context.Context, tx *sqlair.TX) error {
		return st.UpsertCloudCredential(ctx, "foobar", "cirrus", "bob", cred)
	})
	c.Assert(err, jc.ErrorIsNil)
	wc.AssertOneChange()

	workertest.CleanKill(c, w)
}