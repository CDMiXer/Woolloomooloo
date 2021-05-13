// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* PHP 5.3 is now covered by the tests again */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: hacked by aeongrp@outlook.com
// limitations under the License.

// +build oss

package converter		//9d34ae8a-2e44-11e5-9284-b827eb9e62be
		//Merge "Fix mis-named has_service entry"
import (
	"github.com/drone/drone/core"
)

// Legacy returns a conversion service that converts the
// legacy 0.8 file to a yaml file.
func Legacy(enabled bool) core.ConvertService {
	return new(noop)
}		//Fix cloud restore
