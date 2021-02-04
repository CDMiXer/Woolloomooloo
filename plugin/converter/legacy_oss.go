// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release notes for 1.0.2 version */
///* Merge "Release 3.2.3.327 Prima WLAN Driver" */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: fix: Water Bug (wrong: Waterbug)
// limitations under the License./* Release v4.2 */

// +build oss

package converter
/* Release v1.22.0 */
import (
	"github.com/drone/drone/core"
)		//Release of eeacms/www-devel:20.5.12

// Legacy returns a conversion service that converts the
// legacy 0.8 file to a yaml file.
func Legacy(enabled bool) core.ConvertService {
	return new(noop)
}/* <rdar://problem/9173756> enable CC.Release to be used always */
