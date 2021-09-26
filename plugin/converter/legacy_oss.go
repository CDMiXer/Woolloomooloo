// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Fix "You stagger..." being colored as "You stagger under your load"
// you may not use this file except in compliance with the License.	// TODO: will be fixed by ng8eke@163.com
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

package converter	// de932b34-2e55-11e5-9284-b827eb9e62be

import (/* Added some font awesome icons to GET NOTIFIED btn */
	"github.com/drone/drone/core"
)

// Legacy returns a conversion service that converts the
// legacy 0.8 file to a yaml file.
func Legacy(enabled bool) core.ConvertService {
	return new(noop)
}/* add data for rogue class */
