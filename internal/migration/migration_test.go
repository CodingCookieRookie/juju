// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package migration_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/juju/clock"
	"github.com/juju/description/v9"
	"github.com/juju/errors"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/controller"
	"github.com/juju/juju/core/model"
	"github.com/juju/juju/core/modelmigration"
	"github.com/juju/juju/core/resource"
	resourcetesting "github.com/juju/juju/core/resource/testing"
	corestorage "github.com/juju/juju/core/storage"
	domaincharm "github.com/juju/juju/domain/application/charm"
	"github.com/juju/juju/internal/charm"
	loggertesting "github.com/juju/juju/internal/logger/testing"
	"github.com/juju/juju/internal/migration"
	"github.com/juju/juju/internal/storage"
	"github.com/juju/juju/internal/storage/provider"
	"github.com/juju/juju/internal/tools"
	"github.com/juju/juju/internal/version"
	"github.com/juju/juju/state"
)

type ExportSuite struct {
	storageRegistryGetter *MockModelStorageRegistryGetter
	operationsExporter    *MockOperationExporter
	coordinator           *MockCoordinator
	model                 *MockModel
}

var _ = gc.Suite(&ExportSuite{})

func (s *ExportSuite) setupMocks(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.storageRegistryGetter = NewMockModelStorageRegistryGetter(ctrl)
	s.operationsExporter = NewMockOperationExporter(ctrl)
	s.coordinator = NewMockCoordinator(ctrl)
	s.model = NewMockModel(ctrl)

	return ctrl
}

func (s *ExportSuite) TestExportValidates(c *gc.C) {
	defer s.setupMocks(c).Finish()

	scope := modelmigration.NewScope(nil, nil, nil)

	// The order of the expectations is important here. We expect that the
	// validation is the last thing that happens.
	gomock.InOrder(
		s.operationsExporter.EXPECT().ExportOperations(s.storageRegistryGetter),
		s.coordinator.EXPECT().Perform(gomock.Any(), scope, s.model).Return(nil),
		s.model.EXPECT().Validate().Return(nil),
	)

	exporter := migration.NewModelExporter(
		s.operationsExporter,
		nil,
		scope,
		s.storageRegistryGetter,
		s.coordinator,
		nil, nil,
	)

	_, err := exporter.Export(context.Background(), s.model)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *ExportSuite) TestExportValidationFails(c *gc.C) {
	defer s.setupMocks(c).Finish()

	scope := modelmigration.NewScope(nil, nil, nil)

	s.operationsExporter.EXPECT().ExportOperations(s.storageRegistryGetter)
	s.model.EXPECT().Validate().Return(errors.New("boom"))
	s.coordinator.EXPECT().Perform(gomock.Any(), scope, s.model).Return(nil)

	exporter := migration.NewModelExporter(
		s.operationsExporter,
		nil,
		scope,
		s.storageRegistryGetter,
		s.coordinator,
		nil, nil,
	)

	_, err := exporter.Export(context.Background(), s.model)
	c.Assert(err, gc.ErrorMatches, "boom")
}

type ImportSuite struct {
	testing.IsolationSuite
	charmService *MockCharmService
}

var _ = gc.Suite(&ImportSuite{})

func (s *ImportSuite) setupMocks(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.charmService = NewMockCharmService(ctrl)

	return ctrl
}

func (s *ImportSuite) TestBadBytes(c *gc.C) {
	bytes := []byte("not a model")
	scope := func(model.UUID) modelmigration.Scope { return modelmigration.NewScope(nil, nil, nil) }
	controller := &fakeImporter{}
	importer := migration.NewModelImporter(
		controller, scope, nil, nil,
		corestorage.ConstModelStorageRegistry(func() storage.ProviderRegistry {
			return provider.CommonStorageProviders()
		}),
		nil,
		loggertesting.WrapCheckLog(c),
		clock.WallClock,
	)
	model, st, err := importer.ImportModel(context.Background(), bytes)
	c.Check(st, gc.IsNil)
	c.Check(model, gc.IsNil)
	c.Assert(err, gc.ErrorMatches, "yaml: unmarshal errors:\n.*")
}

const modelYaml = `
cloud: dev
config:
  name: foo
  type: lxd
  uuid: bd3fae18-5ea1-4bc5-8837-45400cf1f8f6
actions:
  actions: []
  version: 1
applications:
  applications: []
  version: 1
cloud-image-metadata:
  cloudimagemetadata: []
  version: 1
filesystems:
  filesystems: []
  version: 1
ip-addresses:
  ip-addresses: []
  version: 1
link-layer-devices:
  link-layer-devices: []
  version: 1
machines:
  machines: []
  version: 1
owner: admin
relations:
  relations: []
  version: 1
sequences:
  machine: 2
spaces:
  spaces: []
  version: 1
ssh-host-keys:
  ssh-host-keys: []
  version: 1
storage-pools:
  pools: []
  version: 1
storages:
  storages: []
  version: 1
subnets:
  subnets: []
  version: 1
users:
  users: []
  version: 1
volumes:
  volumes: []
  version: 1
version: 1
`

func (s *ImportSuite) TestUploadBinariesConfigValidate(c *gc.C) {
	type T migration.UploadBinariesConfig // alias for brevity

	check := func(modify func(*T), missing string) {
		config := T{
			CharmService:       struct{ migration.CharmService }{},
			CharmUploader:      struct{ migration.CharmUploader }{},
			ToolsDownloader:    struct{ migration.ToolsDownloader }{},
			ToolsUploader:      struct{ migration.ToolsUploader }{},
			ResourceDownloader: struct{ migration.ResourceDownloader }{},
			ResourceUploader:   struct{ migration.ResourceUploader }{},
		}
		modify(&config)
		realConfig := migration.UploadBinariesConfig(config)
		c.Check(realConfig.Validate(), gc.ErrorMatches, fmt.Sprintf("missing %s not valid", missing))
	}

	check(func(c *T) { c.CharmService = nil }, "CharmService")
	check(func(c *T) { c.CharmUploader = nil }, "CharmUploader")
	check(func(c *T) { c.ToolsDownloader = nil }, "ToolsDownloader")
	check(func(c *T) { c.ToolsUploader = nil }, "ToolsUploader")
	check(func(c *T) { c.ResourceDownloader = nil }, "ResourceDownloader")
	check(func(c *T) { c.ResourceUploader = nil }, "ResourceUploader")
}

func (s *ImportSuite) TestBinariesMigration(c *gc.C) {
	defer s.setupMocks(c).Finish()

	downloader := &fakeDownloader{}
	uploader := &fakeUploader{
		tools:     make(map[version.Binary]string),
		resources: make(map[string]string),
	}

	toolsMap := map[version.Binary]string{
		version.MustParseBinary("2.1.0-ubuntu-amd64"): "/tools/0",
		version.MustParseBinary("2.0.0-ubuntu-amd64"): "/tools/1",
	}

	app0Res := resourcetesting.NewResource(c, nil, "blob0", "app0", "blob0").Resource
	app1Res := resourcetesting.NewResource(c, nil, "blob1", "app1", "blob1").Resource
	app2Res := resourcetesting.NewPlaceholderResource(c, "blob2", "app2")
	resources := []resource.Resource{app0Res, app1Res, app2Res}

	s.charmService.EXPECT().GetCharmArchive(gomock.Any(), domaincharm.CharmLocator{
		Name:     "postgresql",
		Revision: 42,
		Source:   domaincharm.CharmHubSource,
	}).Return(ioutil.NopCloser(strings.NewReader("postgresql content")), "hash0123", nil)
	s.charmService.EXPECT().GetCharmArchive(gomock.Any(), domaincharm.CharmLocator{
		Name:     "magic",
		Revision: 2,
		Source:   domaincharm.LocalSource,
	}).Return(ioutil.NopCloser(strings.NewReader("magic content")), "hash0123", nil)
	s.charmService.EXPECT().GetCharmArchive(gomock.Any(), domaincharm.CharmLocator{
		Name:     "magic",
		Revision: 10,
		Source:   domaincharm.LocalSource,
	}).Return(ioutil.NopCloser(strings.NewReader("magic content")), "hash0123", nil)
	config := migration.UploadBinariesConfig{
		Charms: []string{
			// These 2 are out of order. Rev 2 must be uploaded first.
			"local:trusty/magic-10",
			"local:trusty/magic-2",
			"ch:trusty/postgresql-42",
		},
		CharmService:       s.charmService,
		CharmUploader:      uploader,
		Tools:              toolsMap,
		ToolsDownloader:    downloader,
		ToolsUploader:      uploader,
		Resources:          resources,
		ResourceDownloader: downloader,
		ResourceUploader:   uploader,
	}
	err := migration.UploadBinaries(context.Background(), config, loggertesting.WrapCheckLog(c))
	c.Assert(err, jc.ErrorIsNil)

	expectedCurls := []string{
		// Note ordering.
		"ch:trusty/postgresql-42",
		"local:trusty/magic-2",
		"local:trusty/magic-10",
	}
	c.Assert(uploader.curls, jc.DeepEquals, expectedCurls)

	expectedRefs := []string{
		"postgresql-hash0123",
		"magic-hash0123",
		"magic-hash0123",
	}
	c.Assert(uploader.charmRefs, jc.DeepEquals, expectedRefs)

	c.Assert(downloader.uris, jc.SameContents, []string{
		"/tools/0",
		"/tools/1",
	})
	c.Assert(uploader.tools, jc.DeepEquals, toolsMap)

	c.Assert(downloader.resources, jc.SameContents, []string{
		"app0/blob0",
		"app1/blob1",
	})
	c.Assert(uploader.resources, jc.DeepEquals, map[string]string{
		"app0/blob0": "blob0",
		"app1/blob1": "blob1",
	})
}

func (s *ImportSuite) TestWrongCharmURLAssigned(c *gc.C) {
	defer s.setupMocks(c).Finish()

	downloader := &fakeDownloader{}
	uploader := &fakeUploader{
		reassignCharmURL: true,
	}

	s.charmService.EXPECT().GetCharmArchive(gomock.Any(), domaincharm.CharmLocator{
		Name:     "bar",
		Revision: 2,
		Source:   domaincharm.CharmHubSource,
	}).Return(ioutil.NopCloser(strings.NewReader("bar content")), "hash0123", nil)
	config := migration.UploadBinariesConfig{
		Charms:             []string{"ch:foo/bar-2"},
		CharmService:       s.charmService,
		CharmUploader:      uploader,
		ToolsDownloader:    downloader,
		ToolsUploader:      uploader,
		ResourceDownloader: downloader,
		ResourceUploader:   uploader,
	}
	err := migration.UploadBinaries(context.Background(), config, loggertesting.WrapCheckLog(c))
	c.Assert(err, gc.ErrorMatches,
		"cannot upload charms: charm ch:foo/bar-2 unexpectedly assigned ch:foo/bar-100")
}

type fakeImporter struct {
	model            description.Model
	st               *state.State
	m                *state.Model
	controllerConfig controller.Config
}

func (i *fakeImporter) Import(model description.Model, controllerConfig controller.Config) (*state.Model, *state.State, error) {
	i.model = model
	i.controllerConfig = controllerConfig
	return i.m, i.st, nil
}

type fakeDownloader struct {
	uris      []string
	resources []string
}

func (d *fakeDownloader) OpenURI(_ context.Context, uri string, query url.Values) (io.ReadCloser, error) {
	if query != nil {
		panic("query should be empty")
	}
	d.uris = append(d.uris, uri)
	// Return the URI string as fake content
	return io.NopCloser(bytes.NewReader([]byte(uri))), nil
}

func (d *fakeDownloader) OpenResource(_ context.Context, app, name string) (io.ReadCloser, error) {
	d.resources = append(d.resources, app+"/"+name)
	// Use the resource name as the content.
	return io.NopCloser(bytes.NewReader([]byte(name))), nil
}

type fakeUploader struct {
	tools            map[version.Binary]string
	curls            []string
	charmRefs        []string
	resources        map[string]string
	reassignCharmURL bool
}

func (f *fakeUploader) UploadTools(_ context.Context, r io.Reader, v version.Binary) (tools.List, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, errors.Trace(err)
	}
	f.tools[v] = string(data)
	return tools.List{&tools.Tools{Version: v}}, nil
}

func (f *fakeUploader) UploadCharm(_ context.Context, curl string, charmRef string, r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return "", errors.Trace(err)
	}
	if string(data) != charm.MustParseURL(curl).Name+" content" {
		panic(fmt.Sprintf("unexpected charm body for %s: %s", curl, data))
	}
	f.curls = append(f.curls, curl)
	f.charmRefs = append(f.charmRefs, charmRef)

	outU := curl
	if f.reassignCharmURL {
		outU = charm.MustParseURL(outU).WithRevision(100).String()
	}
	return outU, nil
}

func (f *fakeUploader) UploadResource(_ context.Context, res resource.Resource, r io.Reader) error {
	body, err := io.ReadAll(r)
	if err != nil {
		return errors.Trace(err)
	}
	f.resources[res.ApplicationName+"/"+res.Name] = string(body)
	return nil
}
