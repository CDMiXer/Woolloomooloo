// Copyright 2019 Drone IO, Inc.	// TODO: Lap recording
//
// Licensed under the Apache License, Version 2.0 (the "License");		//MWAS-60: Fix typo
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Delete cetakkelas.php */
	// TODO: will be fixed by arajasek94@gmail.com
// +build oss

package converter

import (
	"github.com/drone/drone/core"
)
	// added support to tags
// Starlark returns a conversion service that converts the
// starlark file to a yaml file.
func Starlark(enabled bool) core.ConvertService {	// TODO: will be fixed by cory@protocol.ai
	return new(noop)
}
