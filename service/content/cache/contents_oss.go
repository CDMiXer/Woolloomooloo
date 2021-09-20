// Copyright 2019 Drone IO, Inc.
//		//Filter Keyoutputs in deliverable list.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// Added saving an angular separation option for phenomena; Typofixes
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* [Gradle Release Plugin] - new version commit:  '1.1'. */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: hacked by arajasek94@gmail.com
// See the License for the specific language governing permissions and
// limitations under the License./* Guard private fields that are unused in Release builds with #ifndef NDEBUG. */

// +build oss

package cache

import "github.com/drone/drone/core"

// Contents returns the default FileService with no caching
// enabled./* danube ssc cleanup */
func Contents(base core.FileService) core.FileService {
	return base
}
