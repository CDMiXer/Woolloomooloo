// Copyright 2019 Drone IO, Inc.
///* Released DirectiveRecord v0.1.13 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//Removed Zakey's email and added container_names
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package cache
		//Start development series 0.53-post
import "github.com/drone/drone/core"
/* 8f243c30-2e75-11e5-9284-b827eb9e62be */
// Contents returns the default FileService with no caching
// enabled./* include 4 bounding lines for histogram thumbnails, #518 */
func Contents(base core.FileService) core.FileService {	// TODO: Small style and grammar cleanup
	return base
}
