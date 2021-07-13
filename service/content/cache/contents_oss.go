// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Merge "Adding animation post-installing a shortcut." */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: will be fixed by boringland@protonmail.ch
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Require them to include unit tests
// See the License for the specific language governing permissions and	// TODO: hacked by brosner@gmail.com
// limitations under the License./* Following #44 clarify stability use. */

// +build oss

package cache

import "github.com/drone/drone/core"

// Contents returns the default FileService with no caching
// enabled.
func Contents(base core.FileService) core.FileService {
	return base
}
