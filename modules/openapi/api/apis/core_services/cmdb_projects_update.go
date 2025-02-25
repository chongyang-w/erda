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

package core_services

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apis"
	"github.com/erda-project/erda/modules/openapi/api/spec"
)

/**
already migration
*/
var CMDB_PROJECT_UPDATE = apis.ApiSpec{
	Path:         "/api/projects/<projectId>",
	BackendPath:  "/api/projects/<projectId>",
	Host:         "core-services.marathon.l4lb.thisdcos.directory:9526",
	Scheme:       "http",
	Method:       "PUT",
	CheckLogin:   true,
	CheckToken:   true,
	IsOpenAPI:    true,
	RequestType:  apistructs.ProjectUpdateRequest{},
	ResponseType: apistructs.ProjectUpdateResponse{},
	Doc:          "summary: 更新项目",
	Audit: func(ctx *spec.AuditContext) error {
		projectID, err := ctx.GetParamInt64("projectId")
		if err != nil {
			return err
		}
		project, err := ctx.GetProject(projectID)
		if err != nil {
			return err
		}
		return ctx.CreateAudit(&apistructs.Audit{
			ScopeType:    apistructs.ProjectScope,
			ScopeID:      uint64(projectID),
			ProjectID:    uint64(projectID),
			TemplateName: apistructs.UpdateProjectTemplate,
			Context:      map[string]interface{}{"projectName": project.Name},
		})
	},
}
