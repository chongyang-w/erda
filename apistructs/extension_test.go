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

package apistructs

import (
	"testing"

	"github.com/alecthomas/assert"
)

func Test_CheckVersion(t *testing.T) {
	spec := Spec{SupportedVersions: []string{">=1.2"}}
	v := spec.CheckDiceVersion("1.3.0-alpha")
	assert.True(t, v)
	v = spec.CheckDiceVersion("1.1-alpha")
	assert.False(t, v)
	v = spec.CheckDiceVersion("1.3.0")
	assert.True(t, v)
	v = spec.CheckDiceVersion("1.1")
	assert.False(t, v)
}
