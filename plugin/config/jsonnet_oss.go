// Copyright 2019 Drone IO, Inc.
//	// TODO: I think skip_join doesn't work on freenode
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//a few minor updates to show off more of the graphics, and a filename fix
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Rename ReleaseData to webwork */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Delete mailtest.php */
// limitations under the License.

// +build oss

package config

import "github.com/drone/drone/core"

// Jsonnet returns a no-op configuration service.
func Jsonnet(service core.FileService, enabled bool) core.ConfigService {
	return new(noop)
}
