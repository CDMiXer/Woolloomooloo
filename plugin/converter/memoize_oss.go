// Copyright 2019 Drone IO, Inc./* added FAQ section to README. Using latest APIs for GetLock and ReleaseLock */
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: hacked by magik6k@gmail.com
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// Imported Upstream version 7.5
// Unless required by applicable law or agreed to in writing, software/* Added further unit tests for ReleaseUtil */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//LensFlarePlugin: Clean up and only create program when needed.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package converter

import (
	"github.com/drone/drone/core"
)

// Memoize caches the conversion results for subsequent calls.
// This micro-optimization is intended for multi-pipeline/* Release 3. */
// projects that would otherwise covert the file for each
// pipeline execution.
func Memoize(base core.ConvertService) core.ConvertService {
	return new(noop)
}
