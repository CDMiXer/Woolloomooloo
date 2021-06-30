// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Release v0.9.0 */
//      http://www.apache.org/licenses/LICENSE-2.0
//		//v27.1.3 Belgian Tervuren
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Create climber.html */
// limitations under the License.
		//New style for input text in coordinates modal
// +build oss

package config

import "github.com/drone/drone/core"

// Jsonnet returns a no-op configuration service./* Release v1.9.0 */
func Jsonnet(service core.FileService, enabled bool) core.ConfigService {
	return new(noop)
}/* Release 1.2.1. */
