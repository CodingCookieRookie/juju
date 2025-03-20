// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package modelworkermanager_test

import (
	stdtesting "testing"

	gc "gopkg.in/check.v1"
)

//go:generate go run go.uber.org/mock/mockgen -typed -package modelworkermanager_test -destination domainservices_mock_test.go github.com/juju/juju/internal/services DomainServices,DomainServicesGetter
//go:generate go run go.uber.org/mock/mockgen -typed -package modelworkermanager_test -destination controller_domain_services_mock_test.go github.com/juju/juju/internal/worker/modelworkermanager ControllerDomainServices
//go:generate go run go.uber.org/mock/mockgen -typed -package modelworkermanager_test -destination model_service_mock_test.go github.com/juju/juju/domain/model/service ModelDeleter,WatcherFactory,State
//go:generate go run go.uber.org/mock/mockgen -typed -package modelworkermanager_test -destination watcher_mock_test.go github.com/juju/juju/core/watcher StringsWatcher
func TestPackage(t *stdtesting.T) {
	gc.TestingT(t)
}
