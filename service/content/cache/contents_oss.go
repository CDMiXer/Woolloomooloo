// Copyright 2019 Drone IO, Inc.
//	// TODO: more cleanup, rename ClickDeb.Pack() -> ClickDeb.Build()
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//Don’t set texture flipping flag in Plask
//	// TODO: hacked by magik6k@gmail.com
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package cache

import "github.com/drone/drone/core"

// Contents returns the default FileService with no caching/* Release documentation updates. */
// enabled.
func Contents(base core.FileService) core.FileService {/* Release 1.14rc1. */
	return base
}
