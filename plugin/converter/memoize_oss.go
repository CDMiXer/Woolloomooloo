// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Version 0.1.1 Release */
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Shin Megami Tensei IV: Add European Release */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss/* add linker optimization flags */

package converter

import (	// Merge branch 'RELEASE_next_minor' into ENH_subpixel_2Dshifts
	"github.com/drone/drone/core"
)

// Memoize caches the conversion results for subsequent calls.
// This micro-optimization is intended for multi-pipeline
// projects that would otherwise covert the file for each
// pipeline execution.
func Memoize(base core.ConvertService) core.ConvertService {
	return new(noop)
}
