// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"context"
	"slices"
	"time"

	"github.com/canonical/sqlair"
	"github.com/juju/clock"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/core/cloudimagemetadata"
	"github.com/juju/juju/domain/cloudimagemetadata/errors"
	schematesting "github.com/juju/juju/domain/schema/testing"
	dberrors "github.com/juju/juju/internal/database"
	loggertesting "github.com/juju/juju/internal/logger/testing"
)

type stateSuite struct {
	schematesting.ControllerSuite

	state *State
}

var _ = gc.Suite(&stateSuite{})

func (s *stateSuite) SetUpTest(c *gc.C) {
	s.ControllerSuite.SetUpTest(c)
	s.state = NewState(s.TxnRunnerFactory(), clock.WallClock, loggertesting.WrapCheckLog(c))
}

// architecture represents a data model for a system architecture with a unique ID and name.
// It is used to prepare and execute sqlair statement.
type architecture struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func (s *stateSuite) TestArchitectureIDsByName(c *gc.C) {
	db, err := s.state.DB()
	c.Assert(err, jc.ErrorIsNil)

	var loadArchsStmt *sqlair.Statement
	loadArchsStmt, err = sqlair.Prepare(`SELECT &architecture.* FROM architecture;`, architecture{})
	c.Assert(err, jc.ErrorIsNil)

	var archs []architecture
	err = db.Txn(context.Background(), func(ctx context.Context, tx *sqlair.TX) error {
		return tx.Query(ctx, loadArchsStmt).GetAll(&archs)
	})
	c.Assert(err, jc.ErrorIsNil)

	obtained := make(map[string]int, len(archs))
	for _, arch := range archs {
		obtained[arch.Name] = arch.ID
	}
	c.Assert(obtained, jc.DeepEquals, architectureIDsByName)
}

// TestSaveMetadata verifies that the metadata is saved correctly in the database
// and checks that creation time is set appropriately.
func (s *stateSuite) TestSaveMetadata(c *gc.C) {
	// Arrange
	testBeginTime := time.Now().Truncate(time.Second) // avoid truncate issue on dqlite creationTime check
	attrs := cloudimagemetadata.MetadataAttributes{
		Stream:          "stream",
		Region:          "region-test",
		Version:         "22.04",
		Arch:            "amd64",
		VirtType:        "virtType-test",
		RootStorageType: "rootStorageType-test",
		Source:          "test",
	}
	expected := []cloudimagemetadata.Metadata{
		{MetadataAttributes: attrs, Priority: 42, ImageId: "42"},
	}

	//  Act
	err := s.state.SaveMetadata(context.Background(), expected)
	c.Assert(err, jc.ErrorIsNil)

	// Assert
	obtained, err := s.retrieveMetadataFromDB()
	for i := range obtained {
		c.Check(obtained[i].CreationTime, jc.After, testBeginTime)
		obtained[i].CreationTime = time.Time{} // ignore time since already checked.
	}
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(obtained, jc.SameContents, expected)
}

// TestSaveMetadataWithDateCreated tests the SaveMetadata method by ensuring metadata is saved with the correct creation date.
func (s *stateSuite) TestSaveMetadataWithDateCreated(c *gc.C) {
	// Arrange
	testBeginTime := time.Now().UTC().Truncate(time.Second) // avoid truncate issue on dqlite creationTime check
	attrs := cloudimagemetadata.MetadataAttributes{
		Stream:          "stream",
		Region:          "region-test",
		Version:         "22.04",
		Arch:            "amd64",
		VirtType:        "virtType-test",
		RootStorageType: "rootStorageType-test",
		Source:          "test",
	}
	expected := []cloudimagemetadata.Metadata{
		{MetadataAttributes: attrs, Priority: 42, ImageId: "42", CreationTime: testBeginTime},
	}

	//  Act
	err := s.state.SaveMetadata(context.Background(), expected)
	c.Assert(err, jc.ErrorIsNil)

	// Assert
	obtained, err := s.retrieveMetadataFromDB()
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(obtained, jc.SameContents, expected)
}

// TestSaveMetadataSeveralMetadata verifies that multiple metadata entries are saved correctly in the database.
func (s *stateSuite) TestSaveMetadataSeveralMetadata(c *gc.C) {
	// Arrange
	testBeginTime := time.Now().Truncate(time.Second) // avoid truncate issue on dqlite creationTime check
	attrs1 := cloudimagemetadata.MetadataAttributes{
		Stream:          "stream",
		Region:          "region-test",
		Version:         "22.04",
		Arch:            "arm64",
		VirtType:        "virtType-test",
		RootStorageType: "rootStorageType-test",
		Source:          "test",
	}
	attrs2 := cloudimagemetadata.MetadataAttributes{
		Stream:          "chalk",
		Region:          "nether",
		Version:         "12.04",
		Arch:            "amd64",
		Source:          "test",
		RootStorageSize: ptr(uint64(1024)),
	}
	expected := []cloudimagemetadata.Metadata{
		{MetadataAttributes: attrs1, ImageId: "1"},
		{MetadataAttributes: attrs2, ImageId: "2"},
	}

	//  Act
	err := s.state.SaveMetadata(context.Background(), expected)
	c.Assert(err, jc.ErrorIsNil)

	// Assert
	obtained, err := s.retrieveMetadataFromDB()
	for i := range obtained {
		c.Check(obtained[i].CreationTime, jc.After, testBeginTime)
		obtained[i].CreationTime = time.Time{} // ignore time since already checked.
	}
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(obtained, jc.SameContents, expected)
}

func (s *stateSuite) TestSaveMetadataUpdateMetadata(c *gc.C) {
	// Arrange
	testBeginTime := time.Now().Truncate(time.Second) // avoid truncate issue on dqlite creationTime check
	attrs1 := cloudimagemetadata.MetadataAttributes{
		Stream:          "stream",
		Region:          "region-test",
		Version:         "22.04",
		Arch:            "arm64",
		VirtType:        "virtType-test",
		RootStorageType: "rootStorageType-test",
		Source:          "test",
	}
	attrs2 := attrs1
	attrs2.RootStorageSize = ptr(uint64(1024)) // Not part of the key, but shouldn't be updated either

	//  Act
	err := s.state.SaveMetadata(context.Background(), []cloudimagemetadata.Metadata{
		{MetadataAttributes: attrs1, ImageId: "1"},
	})
	c.Assert(err, jc.ErrorIsNil)
	err = s.state.SaveMetadata(context.Background(), []cloudimagemetadata.Metadata{
		{MetadataAttributes: attrs2, ImageId: "2"},
	})
	c.Assert(err, jc.ErrorIsNil)

	// Assert
	obtained, err := s.retrieveMetadataFromDB()
	for i := range obtained {
		c.Check(obtained[i].CreationTime, jc.After, testBeginTime)
		obtained[i].CreationTime = time.Time{} // ignore time since already checked.
	}
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(obtained, jc.SameContents, []cloudimagemetadata.Metadata{
		{MetadataAttributes: attrs1, ImageId: "2"}, // Imageid has been updated, but other attributes don't.
	})
}

func (s *stateSuite) TestSaveMetadataWithSameAttributes(c *gc.C) {
	// Arrange
	attrs1 := cloudimagemetadata.MetadataAttributes{
		Stream:          "stream",
		Region:          "region-test",
		Version:         "22.04",
		Arch:            "arm64",
		VirtType:        "virtType-test",
		RootStorageType: "rootStorageType-test",
		Source:          "test",
	}
	attrs2 := attrs1
	attrs2.RootStorageSize = ptr(uint64(1024)) // Not part of the key, but shouldn't be updated either

	//  Act
	err := s.state.SaveMetadata(context.Background(), []cloudimagemetadata.Metadata{
		{MetadataAttributes: attrs1, ImageId: "1"},
		{MetadataAttributes: attrs2, ImageId: "2"},
	})

	// Assert
	c.Assert(dberrors.IsErrConstraintUnique(err), jc.IsTrue)
}

// TestSaveMetadataSeveralMetadataWithInvalidArchitecture verifies that metadata with an invalid architecture is ignored
// when saving multiple metadata entries.
func (s *stateSuite) TestSaveMetadataSeveralMetadataWithInvalidArchitecture(c *gc.C) {
	// Arrange
	attrsBase := cloudimagemetadata.MetadataAttributes{
		Stream:          "stream",
		Region:          "region-test",
		Version:         "22.04",
		Arch:            "arm64",
		VirtType:        "virtType-test",
		RootStorageType: "rootStorageType-test",
		Source:          "test",
	}
	attr1 := attrsBase
	attr2 := attrsBase
	attrIncorrectArch := attrsBase

	attr2.Region = "anotherRegion"
	attrIncorrectArch.Arch = "unknownArch"

	inserted := []cloudimagemetadata.Metadata{
		{MetadataAttributes: attr1, ImageId: "1"},
		{MetadataAttributes: attrIncorrectArch, ImageId: "2"},
		{MetadataAttributes: attr2, ImageId: "3"},
	}

	//  Act
	err := s.state.SaveMetadata(context.Background(), inserted)

	// Assert
	c.Assert(err, gc.ErrorMatches, ".*unknown architecture.*")
}

// TestDeleteMetadataWithImageID verifies that the DeleteMetadataWithImageID method correctly removes specified entries from the cloud_image_metadata table.
func (s *stateSuite) TestDeleteMetadataWithImageID(c *gc.C) {
	// Arrange
	s.runQuery(`
INSERT INTO cloud_image_metadata (uuid, created_at,source,stream,region,version,architecture_id,virt_type,root_storage_type,priority,image_id)
VALUES 
('a', datetime('now','localtime'), 'custom', 'stream', 'region-1', '22.04',0, 'virtType-test', 'rootStorageType-test', 42, 'to-keep'),
('b', datetime('now','localtime'), 'custom', 'stream', 'region-2', '22.04',0, 'virtType-test', 'rootStorageType-test', 3, 'to-delete'),
('c', datetime('now','localtime'), 'custom', 'stream', 'region-3', '22.04',0, 'virtType-test', 'rootStorageType-test', 42, 'to-delete')
`)

	//  Act
	err := s.state.DeleteMetadataWithImageID(context.Background(), "to-delete")

	// Assert
	c.Assert(err, jc.ErrorIsNil)
	obtained, err := s.retrieveMetadataFromDB()
	for i := range obtained {
		obtained[i].CreationTime = time.Time{} // ignore time
	}
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(obtained, jc.SameContents, []cloudimagemetadata.Metadata{
		{MetadataAttributes: cloudimagemetadata.MetadataAttributes{
			Stream:          "stream",
			Region:          "region-1",
			Version:         "22.04",
			Arch:            "amd64",
			VirtType:        "virtType-test",
			RootStorageType: "rootStorageType-test",
			Source:          "custom",
		}, Priority: 42, ImageId: "to-keep"},
	})
}

// TestFindMetadata tests the retrieval of metadata from the cloud_image_metadata table using various filters.
func (s *stateSuite) TestFindMetadata(c *gc.C) {
	// Arrange
	err := s.runQuery(`
INSERT INTO cloud_image_metadata (uuid,created_at,source,stream,region,version,architecture_id,virt_type,root_storage_type,priority,image_id)
VALUES 
('a',datetime('now','localtime'), 'custom', 'stream', 'region', '08.04',0, 'virtType', 'storage', 42, 'id'),
('b',datetime('now','localtime'), 'custom', 'unique', 'region', '10.04',1, 'virtType', 'storage', 42, 'id'),
('c',datetime('now','localtime'), 'custom', 'stream', 'unique', '12.04',2, 'virtType', 'storage', 42, 'id'),
('d',datetime('now','localtime'), 'custom', 'stream', 'region', '14.00',3, 'virtType', 'storage', 42, 'id'),
('e',datetime('now','localtime'), 'custom', 'stream', 'region', '16.04',4, 'virtType', 'storage', 42, 'id'),
('f',datetime('now','localtime'), 'custom', 'stream', 'region', '18.04',0, 'unique', 'storage', 42, 'id'),
('g',datetime('now','localtime'), 'custom', 'stream', 'region', '20.04',1, 'virtType', 'unique', 42, 'id'),
('h',datetime('now','localtime'), 'custom', 'stream', 'region', '22.04',2, 'virtType', 'storage', 1, 'id'),
('i',datetime('now','localtime'), 'custom', 'stream', 'region', '24.04',3, 'virtType', 'storage', 42, 'unique');
`)
	c.Assert(err, jc.ErrorIsNil)
	expectedBase, err := s.retrieveMetadataFromDB()
	c.Assert(err, jc.ErrorIsNil)

	for _, testCase := range []struct {
		description string
		filter      cloudimagemetadata.MetadataFilter
		accept      func(cloudimagemetadata.Metadata) bool
	}{
		{
			description: "Filter by region: 'unique'",
			filter:      cloudimagemetadata.MetadataFilter{Region: "unique"},
			accept:      func(metadata cloudimagemetadata.Metadata) bool { return metadata.Region == "unique" },
		},
		{
			description: "Filter by version: 'any of [22.04 24.04]'",
			filter:      cloudimagemetadata.MetadataFilter{Versions: []string{"22.04", "24.04"}},
			accept: func(metadata cloudimagemetadata.Metadata) bool {
				return slices.Contains([]string{"22.04", "24.04"}, metadata.Version)
			},
		},
		{
			description: "Filter by architecture: 'any of [amd64(id:0) arm64(id:1)]'",
			filter:      cloudimagemetadata.MetadataFilter{Arches: []string{"amd64", "arm64"}},
			accept: func(metadata cloudimagemetadata.Metadata) bool {
				return slices.Contains([]string{"amd64", "arm64"}, metadata.Arch)
			},
		},
		{
			description: "Filter by stream: 'unique'",
			filter:      cloudimagemetadata.MetadataFilter{Stream: "unique"},
			accept:      func(metadata cloudimagemetadata.Metadata) bool { return metadata.Stream == "unique" },
		},
		{
			description: "Filter by virt_type: 'unique'",
			filter:      cloudimagemetadata.MetadataFilter{VirtType: "unique"},
			accept:      func(metadata cloudimagemetadata.Metadata) bool { return metadata.VirtType == "unique" },
		},
		{
			description: "Filter by root_storage_type: 'unique'",
			filter:      cloudimagemetadata.MetadataFilter{RootStorageType: "unique"},
			accept:      func(metadata cloudimagemetadata.Metadata) bool { return metadata.RootStorageType == "unique" },
		},
		{
			description: "Filter by root_storage_type: 'storage' and region: 'region' and  stream: 'stream'",
			filter:      cloudimagemetadata.MetadataFilter{RootStorageType: "storage", Region: "region", Stream: "stream"},
			accept: func(metadata cloudimagemetadata.Metadata) bool {
				return metadata.RootStorageType == "storage" && metadata.Region == "region" && metadata.Stream == "stream"
			},
		},
	} {

		// Act
		obtained, err := s.state.FindMetadata(context.Background(), testCase.filter)
		c.Check(err, jc.ErrorIsNil, gc.Commentf("test: %s\n - unexpected error: %s", testCase.description, err))

		// Assert
		c.Check(obtained, jc.SameContents, filter(expectedBase, testCase.accept), gc.Commentf("test: %s\n - obtained value mismatched", testCase.description))
	}
}

// TestFindMetadataNotFound verifies that FindMetadata returns a NotFound error when no matching metadata is found.
func (s *stateSuite) TestFindMetadataNotFound(c *gc.C) {
	// Arrange
	err := s.runQuery(`
INSERT INTO cloud_image_metadata (uuid,created_at,source,stream,region,version,architecture_id,virt_type,root_storage_type,priority,image_id)
VALUES 
('a',datetime('now','localtime'), 'custom', 'stream', 'unique', '08.04',0, 'virtType', 'storage', 42, 'id'),
('b',datetime('now','localtime'), 'custom', 'unique', 'region', '10.04',1, 'virtType', 'storage', 42, 'id');
`)
	c.Assert(err, jc.ErrorIsNil)

	// Act
	_, err = s.state.FindMetadata(context.Background(), cloudimagemetadata.MetadataFilter{Stream: "unique", Region: "unique"})

	// Assert
	c.Assert(err, jc.ErrorIs, errors.NotFound)
}

// TestFindMetadataExpired checks that expired metadata entries are correctly excluded from query results.
func (s *stateSuite) TestFindMetadataExpired(c *gc.C) {
	// Arrange
	err := s.runQuery(`
INSERT INTO cloud_image_metadata (uuid,created_at,source,stream,region,version,architecture_id,virt_type,root_storage_type,priority,image_id)
VALUES 
('a',datetime('now','-3 days'), 'custom', 'stream', 'region', '08.04',0, 'virtType', 'storage', 42, 'id'),
('b',datetime('now','localtime'), 'custom', 'stream', 'region', '10.04',1, 'virtType', 'storage', 42, 'id'),
('c',datetime('now','-3 days'), 'custom', 'stream', 'region', '12.04',1, 'virtType', 'storage', 42, 'id');
`)
	c.Assert(err, jc.ErrorIsNil)

	// Act
	obtained, err := s.state.FindMetadata(context.Background(), cloudimagemetadata.MetadataFilter{})

	// Assert
	c.Assert(err, jc.ErrorIsNil)
	for i := range obtained {
		obtained[i].CreationTime = time.Time{} // ignore time
	}
	c.Assert(obtained, jc.SameContents, []cloudimagemetadata.Metadata{{
		MetadataAttributes: cloudimagemetadata.MetadataAttributes{
			Stream:          "stream",
			Region:          "region",
			Version:         "10.04",
			Arch:            "arm64",
			VirtType:        "virtType",
			RootStorageType: "storage",
			Source:          "custom",
		},
		Priority: 42,
		ImageId:  "id",
	},
	})
}

// TestAllCloudImageMetadata tests the retrieval of all cloud image metadata from the database,
// except expired one.
func (s *stateSuite) TestAllCloudImageMetadata(c *gc.C) {
	// Arrange
	err := s.runQuery(`
INSERT INTO cloud_image_metadata (uuid,created_at,source,stream,region,version,architecture_id,virt_type,root_storage_type,priority,image_id)
VALUES 
('a',datetime('now','-3 days'), 'custom', 'stream', 'region', '08.04',0, 'virtType', 'storage', 42, 'id'),
('b',datetime('now','localtime'), 'custom', 'stream', 'region', '10.04',1, 'virtType', 'storage', 42, 'id'),
('c',datetime('now','-3 days'), 'custom', 'stream', 'region', '12.04',1, 'virtType', 'storage', 42, 'id'),
('d',datetime('now','localtime'), 'custom', 'stream', 'region', '16.04',1, 'virtType', 'storage', 42, 'id');
`)
	c.Assert(err, jc.ErrorIsNil)

	// Act
	obtained, err := s.state.AllCloudImageMetadata(context.Background())

	// Assert
	c.Assert(err, jc.ErrorIsNil)
	for i := range obtained {
		obtained[i].CreationTime = time.Time{} // ignore time
	}
	c.Assert(obtained, jc.SameContents, []cloudimagemetadata.Metadata{
		{
			MetadataAttributes: cloudimagemetadata.MetadataAttributes{
				Stream:          "stream",
				Region:          "region",
				Version:         "10.04",
				Arch:            "arm64",
				VirtType:        "virtType",
				RootStorageType: "storage",
				Source:          "custom",
			},
			Priority: 42,
			ImageId:  "id",
		},
		{
			MetadataAttributes: cloudimagemetadata.MetadataAttributes{
				Stream:          "stream",
				Region:          "region",
				Version:         "16.04",
				Arch:            "arm64",
				VirtType:        "virtType",
				RootStorageType: "storage",
				Source:          "custom",
			},
			Priority: 42,
			ImageId:  "id",
		},
	})
}

// TestAllCloudImageMetadataNoMetadata ensures that AllCloudImageMetadata returns no metadata
// without error if all entries have expired.
func (s *stateSuite) TestAllCloudImageMetadataNoMetadata(c *gc.C) {
	// Arrange
	err := s.runQuery(`
INSERT INTO cloud_image_metadata (uuid,created_at,source,stream,region,version,architecture_id,virt_type,root_storage_type,priority,image_id)
VALUES 
('a',datetime('now','-3 days'), 'custom', 'stream', 'region', '08.04',0, 'virtType', 'storage', 42, 'id'),
('b',datetime('now','-3 days'), 'custom', 'stream', 'region', '12.04',1, 'virtType', 'storage', 42, 'id');
`)
	c.Assert(err, jc.ErrorIsNil)

	// Act
	obtained, err := s.state.AllCloudImageMetadata(context.Background())

	// Assert
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(obtained, jc.SameContents, []cloudimagemetadata.Metadata{}) // No non expired metadata
}

// TestCleanupMetatada verifies that the metadata is properly cleaned up on a new insert in the
// database.
// Custom source never expires.
func (s *stateSuite) TestCleanupMetatada(c *gc.C) {
	// Arrange
	err := s.runQuery(`
INSERT INTO cloud_image_metadata (uuid,created_at,source,stream,region,version,architecture_id,virt_type,root_storage_type,root_storage_size,priority,image_id)
VALUES 
('a',datetime('now','-3 days'), 'cloud', 'stream', 'region', '08.04',0, 'virtType', 'storage', 128, 42, 'id'),
('b',datetime('now','-3 days'), 'custom', 'stream', 'region', '12.04',1, 'virtType', 'storage', 1024, 42, 'id');
`)
	c.Assert(err, jc.ErrorIsNil)
	attrs := cloudimagemetadata.MetadataAttributes{
		Stream:          "stream",
		Region:          "region-test",
		Version:         "22.04",
		Arch:            "amd64",
		VirtType:        "virtType-test",
		RootStorageType: "rootStorageType-test",
		Source:          "test",
	}
	expected := []cloudimagemetadata.Metadata{
		{MetadataAttributes: attrs, Priority: 42, ImageId: "42"},
	}

	//  Act
	err = s.state.SaveMetadata(context.Background(), expected)
	c.Assert(err, jc.ErrorIsNil)

	// Assert
	obtained, err := s.retrieveMetadataFromDB()
	for i := range obtained {
		obtained[i].CreationTime = time.Time{} // ignore time for simplicity
	}
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(obtained, jc.SameContents, append(expected,
		// Custom that has not expired
		cloudimagemetadata.Metadata{MetadataAttributes: cloudimagemetadata.MetadataAttributes{
			Stream:          "stream",
			Region:          "region",
			Version:         "12.04",
			Arch:            "arm64",
			VirtType:        "virtType",
			RootStorageType: "storage",
			RootStorageSize: ptr(uint64(1024)),
			Source:          "custom",
		}, Priority: 42, ImageId: "id"}))
}