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

package filter

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda-infra/providers/component-protocol/utils/cputil"
	"github.com/erda-project/erda/modules/dop/component-protocol/types"
	autotestv2 "github.com/erda-project/erda/modules/dop/services/autotest_v2"
	"github.com/erda-project/erda/modules/dop/services/iteration"
	"github.com/erda-project/erda/modules/dop/services/testcase"
	mttestplan "github.com/erda-project/erda/modules/dop/services/testplan"
	"github.com/erda-project/erda/modules/openapi/component-protocol/components/base"
	"github.com/erda-project/erda/modules/openapi/component-protocol/components/filter"
)

type Filter struct {
	base.DefaultProvider

	sdk *cptype.SDK

	// issue manage
	iteration *iteration.Iteration
	// manual test
	mtTestCase *testcase.Service
	mtTestPlan *mttestplan.TestPlan
	// auto test
	atTestPlan *autotestv2.Service

	InParams InParams `json:"-"`
	State    State    `json:"state,omitempty"`
}

type InParams struct {
	FrontEndProjectID string `json:"projectID,omitempty"`
	ProjectID         uint64
}

type State struct {
	Conditions []filter.PropCondition `json:"conditions,omitempty"`
	Values     SelectedValues         `json:"values,omitempty"`
}

type SelectedValues struct {
	IterationIDs []uint64 `json:"iteration,omitempty"`
}

func (f *Filter) initFromProtocol(ctx context.Context, c *cptype.Component) error {
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, f); err != nil {
		return err
	}

	// sdk
	f.sdk = cputil.SDK(ctx)
	f.iteration = ctx.Value(types.IterationService).(*iteration.Iteration)
	f.mtTestCase = ctx.Value(types.ManualTestCaseService).(*testcase.Service)
	f.mtTestPlan = ctx.Value(types.ManualTestPlanService).(*mttestplan.TestPlan)
	f.atTestPlan = ctx.Value(types.AutoTestPlanService).(*autotestv2.Service)

	// in params
	if err := f.setInParams(ctx); err != nil {
		return err
	}

	return nil
}

func (f *Filter) setInParams(ctx context.Context) error {
	b, err := json.Marshal(cputil.SDK(ctx).InParams)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, &f.InParams); err != nil {
		return err
	}

	f.InParams.ProjectID, err = strconv.ParseUint(f.InParams.FrontEndProjectID, 10, 64)
	return err
}

func (f *Filter) setToComponent(c *cptype.Component) error {
	b, err := json.Marshal(f)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, &c); err != nil {
		return err
	}
	return nil
}
