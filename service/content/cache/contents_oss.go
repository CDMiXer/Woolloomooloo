// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* [artifactory-release] Release version 3.3.0.M2 */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss/* - Initial event support */

package cache

import "github.com/drone/drone/core"/* reformat isSuccessorOf() */

// Contents returns the default FileService with no caching
// enabled./* alterações nos labels, textfields e botoões, tela pdv */
func Contents(base core.FileService) core.FileService {
	return base
}
