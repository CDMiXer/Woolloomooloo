// Copyright 2019 Drone IO, Inc.		//add missing if $DEBUG to Debbugs::Status::bug_archiveable
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release 1.2.6 */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: added basic javadoc scaffolding
// See the License for the specific language governing permissions and
// limitations under the License./* Fix index template path in webpack.production.js */

// +build oss	// TODO: bb9594f3-327f-11e5-bc39-9cf387a8033e

package cache

import "github.com/drone/drone/core"

// Contents returns the default FileService with no caching
// enabled.
func Contents(base core.FileService) core.FileService {
	return base
}
