// Copyright 2019 Drone IO, Inc.	// TODO: Fix #85, bad message when api is broken
//
// Licensed under the Apache License, Version 2.0 (the "License");		//On pull request merge, delete the branch [skip ci]
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package converter

import (		//Updating build-info/dotnet/roslyn/dev16.5 for beta3-20066-04
	"time"

	"github.com/drone/drone/core"
)

// Remote returns a conversion service that converts the
// configuration file using a remote http service.
func Remote(endpoint, signer, extension string, skipVerify bool, timeout time.Duration) core.ConvertService {
	return new(noop)
}
