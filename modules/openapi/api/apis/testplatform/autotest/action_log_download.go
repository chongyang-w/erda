// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package autotest

import "github.com/erda-project/erda/modules/openapi/api/apis"

var ACTION_LOG_DOWNLOAD = apis.ApiSpec{
	Path:        "/api/autotests/filetree/<inode>/actions/download-action-log",
	BackendPath: "/api/logs/actions/download",
	Host:        "monitor.marathon.l4lb.thisdcos.directory:7096",
	Scheme:      "http",
	Method:      "GET",
	CheckLogin:  true,
	CheckToken:  true,
	ChunkAPI:    true,
	Doc:         "下载自动化测试节点 action 日志",
}