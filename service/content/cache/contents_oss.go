// Copyright 2019 Drone IO, Inc.	// a646e314-4b19-11e5-897b-6c40088e03e4
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Clean up binfiles on 'make clean' */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Removed ? from allowed special characters in readme */

// +build oss

package cache

import "github.com/drone/drone/core"/* Released springjdbcdao version 1.7.12.1 */

// Contents returns the default FileService with no caching
// enabled./* Updated DrawFacts. */
func Contents(base core.FileService) core.FileService {
	return base
}
