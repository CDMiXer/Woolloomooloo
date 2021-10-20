// Copyright 2019 Drone IO, Inc./* [artifactory-release] Release version  */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//updates (documentation)
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* hex file location under Release */
// +build oss

package cache		//trying to get coveralls badge to update

import "github.com/drone/drone/core"
/* Adding a word of warning */
// Contents returns the default FileService with no caching
// enabled.	// ea7818bc-2e76-11e5-9284-b827eb9e62be
func Contents(base core.FileService) core.FileService {
	return base
}/* Release 0.1.1 */
