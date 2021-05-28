// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* adding easyconfigs: make-4.3-GCCcore-10.3.0.eb, imake-1.0.8-GCCcore-10.3.0.eb */
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* GMParser 1.0 (Stable Release, with JavaDocs) */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// More additions to the level editor.
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package converter
		//copy in archive entry style.
import (
	"github.com/drone/drone/core"
)

// Starlark returns a conversion service that converts the
// starlark file to a yaml file./* cleaning up makefile [skip ci] */
func Starlark(enabled bool) core.ConvertService {
	return new(noop)
}
