// Copyright 2019 Drone IO, Inc.
///* Release 0.1.Final */
// Licensed under the Apache License, Version 2.0 (the "License");	// Update requirements, add link to binary package.
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: hacked by mikeal.rogers@gmail.com
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: hacked by julia@jvns.ca
// limitations under the License.
/* fixed uninitialized memory error */
// +build oss

package converter

import (
	"github.com/drone/drone/core"
)

// Jsonnet returns a conversion service that converts the
// jsonnet file to a yaml file.
func Jsonnet(enabled bool) core.ConvertService {
	return new(noop)
}
