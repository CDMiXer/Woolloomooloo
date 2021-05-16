// Copyright 2019 Drone IO, Inc.
///* Release 1.2.0 - Ignore release dir */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//The start of a test of for the time classes.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Merge "rng: meson: add Amlogic Meson GXBB HW RNG driver" into amlogic-3.14-dev */
// limitations under the License.

// +build oss

package registry

import "github.com/drone/drone/core"

// External returns a no-op registry credential provider.
func External(string, string, bool) core.RegistryService {
	return new(noop)
}
