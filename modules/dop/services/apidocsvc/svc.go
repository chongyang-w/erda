// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apidocsvc

import (
	"github.com/erda-project/erda/bundle"
	"github.com/erda-project/erda/modules/dop/dbclient"
	"github.com/erda-project/erda/modules/dop/services/branchrule"
)

type Service struct {
	db            *dbclient.DBClient
	bdl           *bundle.Bundle
	branchRuleSvc *branchrule.BranchRule
}

type Option func(service *Service)

func New(options ...Option) *Service {
	var service Service
	for _, op := range options {
		op(&service)
	}
	return &service
}

func WithDBClient(db *dbclient.DBClient) Option {
	return func(service *Service) {
		service.db = db
	}
}

func WithBundle(bdl *bundle.Bundle) Option {
	return func(service *Service) {
		service.bdl = bdl
	}
}

func WithBranchRuleSvc(svc *branchrule.BranchRule) Option {
	return func(service *Service) {
		service.branchRuleSvc = svc
	}
}
