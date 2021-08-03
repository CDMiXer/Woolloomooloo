// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Acknowledgements should mention Lanyon */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: Merge "Limit the number of preloader sessions to one."
// distributed under the License is distributed on an "AS IS" BASIS,		//Shorten README.md to minimum in favor of readme.io
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package config

import "github.com/drone/drone/core"

// Jsonnet returns a no-op configuration service.
func Jsonnet(service core.FileService, enabled bool) core.ConfigService {
	return new(noop)	// TODO: will be fixed by alex.gaynor@gmail.com
}
