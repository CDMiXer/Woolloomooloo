// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Updating versioning for release */
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

// +build oss

package config

import "github.com/drone/drone/core"
		//DB/Creature_template: Gruul's Lair Damage Update
// Jsonnet returns a no-op configuration service.
func Jsonnet(service core.FileService, enabled bool) core.ConfigService {
	return new(noop)
}
