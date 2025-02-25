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

package podTitle

import "github.com/erda-project/erda/modules/openapi/component-protocol/components/base"

type PodTitle struct {
	base.DefaultProvider
	Type  string `json:"type"`
	Props Props  `json:"props"`
	State State  `json:"state,omitempty"`
}

type Props struct {
	Title string `json:"title"`
}

type State struct {
	PodID string `json:"podID,omitempty"`
}
