// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"context"

	jc "github.com/juju/testing/checkers"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	coresecrets "github.com/juju/juju/core/secrets"
	domainsecret "github.com/juju/juju/domain/secret"
	domaintesting "github.com/juju/juju/domain/testing"
	"github.com/juju/juju/internal/secrets/provider"
	"github.com/juju/juju/internal/uuid"
)

func (s *serviceSuite) TestDeleteObsoleteUserSecretRevisions(c *gc.C) {
	ctrl := s.setupMocks(c)
	defer ctrl.Finish()

	revisionID1, err := uuid.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	revisionID2, err := uuid.NewUUID()
	c.Assert(err, jc.ErrorIsNil)

	s.state.EXPECT().DeleteObsoleteUserSecretRevisions(gomock.Any()).Return([]string{revisionID1.String(), revisionID2.String()}, nil)
	s.secretBackendState.EXPECT().RemoveSecretBackendReference(gomock.Any(), revisionID1.String(), revisionID2.String()).Return(nil)

	err = s.service.DeleteObsoleteUserSecretRevisions(context.Background())
	c.Assert(err, jc.ErrorIsNil)
}

func (s *serviceSuite) TestDeleteSecret(c *gc.C) {
	ctrl := s.setupMocks(c)
	defer ctrl.Finish()

	uri := coresecrets.NewURI()

	s.state.EXPECT().GetSecretAccess(gomock.Any(), uri, domainsecret.AccessParams{
		SubjectTypeID: domainsecret.SubjectUnit,
		SubjectID:     "mariadb/0",
	}).Return("manage", nil)
	s.state.EXPECT().DeleteSecret(domaintesting.IsAtomicContextChecker, uri, []int{1, 2})

	revs := provider.SecretRevisions{}
	revs.Add(uri, "rev-id1")
	revs.Add(uri, "rev-id2")

	err := s.service.DeleteSecret(context.Background(), uri, DeleteSecretParams{
		LeaderToken: successfulToken{},
		Accessor: SecretAccessor{
			Kind: UnitAccessor,
			ID:   "mariadb/0",
		},
		Revisions: []int{1, 2},
	})
	c.Assert(err, jc.ErrorIsNil)
}
