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

package addPodFileEditor

import (
	"context"

	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda/modules/cmp"
	"github.com/erda-project/erda/modules/openapi/component-protocol/components/base"
)

type ComponentAddPodFileEditor struct {
	base.DefaultProvider

	sdk        *cptype.SDK
	ctx        context.Context
	server     cmp.SteveServer
	Type       string                 `json:"type"`
	Props      Props                  `json:"props"`
	State      State                  `json:"state"`
	Operations map[string]interface{} `json:"operations,omitempty"`
}

type Props struct {
	Bordered     bool     `json:"bordered"`
	FileValidate []string `json:"fileValidate,omitempty"`
	MinLines     int      `json:"minLines"`
}

type State struct {
	ClusterName string `json:"clusterName,omitempty"`
	Values      Values `json:"values,omitempty"`
	Value       string `json:"value,omitempty"`
}

type Operation struct {
	Key    string `json:"key,omitempty"`
	Reload bool   `json:"reload"`
}

type Values struct {
	Namespace string `json:"namespace,omitempty"`
}
