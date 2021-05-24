// Copyright 2019 Drone IO, Inc.
///* New translations authorization.php (Chinese Simplified) */
// Licensed under the Apache License, Version 2.0 (the "License");		//eb18fd62-2e4f-11e5-9284-b827eb9e62be
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Debug: Add somme logging.
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Increase test timeouts */
// See the License for the specific language governing permissions and	// TODO: will be fixed by vyzo@hackzen.org
// limitations under the License.

// +build oss

package cache

import "github.com/drone/drone/core"

// Contents returns the default FileService with no caching
// enabled.
func Contents(base core.FileService) core.FileService {
	return base
}/* Merge branch 'master' into financial_assessmentws */
