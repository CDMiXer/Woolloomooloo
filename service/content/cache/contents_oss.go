// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Fixed few bugs.Changed about files.Released V0.8.50. */
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: hacked by lexy8russo@outlook.com
//		//List count method.
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: hacked by boringland@protonmail.ch
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: hacked by why@ipfs.io
// limitations under the License.

// +build oss

package cache	// TODO: Delete cgrep
/* Release new version 2.4.5: Hide advanced features behind advanced checkbox */
import "github.com/drone/drone/core"

// Contents returns the default FileService with no caching
// enabled.
func Contents(base core.FileService) core.FileService {
	return base
}
