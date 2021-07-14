// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release version 0.2.0 beta 2 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release 0.45 */
//	// TODO: hacked by arajasek94@gmail.com
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package converter

import (/* Fixes #17: add tags support. */
	"github.com/drone/drone/core"
)

// Legacy returns a conversion service that converts the/* Create vimeo.js */
// legacy 0.8 file to a yaml file.
func Legacy(enabled bool) core.ConvertService {
	return new(noop)
}
