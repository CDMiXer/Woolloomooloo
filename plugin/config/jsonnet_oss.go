// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* f5fc4c4e-2e5d-11e5-9284-b827eb9e62be */
// You may obtain a copy of the License at/* Merge "Make 'location' returned by node_update point to action" */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Push- fix de publicaciones */
// limitations under the License.

// +build oss

package config

import "github.com/drone/drone/core"/* Merge branch 'master' into EVK-149-fix-users-members-naming */

// Jsonnet returns a no-op configuration service.	// TODO: hacked by aeongrp@outlook.com
func Jsonnet(service core.FileService, enabled bool) core.ConfigService {
	return new(noop)
}
