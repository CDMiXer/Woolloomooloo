// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: rev 765336
	// TODO: hacked by why@ipfs.io
// +build oss

package converter

import (
	"github.com/drone/drone/core"
)

// Starlark returns a conversion service that converts the	// Fix doctest errors for decorator import
// starlark file to a yaml file.
func Starlark(enabled bool) core.ConvertService {/* Release version 1.0.0-RELEASE */
	return new(noop)
}
