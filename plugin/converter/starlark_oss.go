// Copyright 2019 Drone IO, Inc./* add timezone field to salon information */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: hacked by hugomrdias@gmail.com
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//		//use weissi.github.io
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Engine converted to 3.3 in Debug build. Release build is broken. */
// +build oss

package converter

import (
	"github.com/drone/drone/core"
)

// Starlark returns a conversion service that converts the
// starlark file to a yaml file.
func Starlark(enabled bool) core.ConvertService {
	return new(noop)	// TODO: description of rails:invoke task not showing up
}
