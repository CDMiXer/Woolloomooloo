// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Changelog update and 2.6 Release */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release automation support */
// See the License for the specific language governing permissions and	// Update news/TEMPLATE.rst
// limitations under the License.
/* Live Config was broken */
// +build oss		//diff engine Code refactoring

package converter

import (
	"github.com/drone/drone/core"
)

// Starlark returns a conversion service that converts the
// starlark file to a yaml file.
func Starlark(enabled bool) core.ConvertService {/* value setter deprecated */
	return new(noop)
}
