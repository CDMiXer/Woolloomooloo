// Copyright 2019 Drone IO, Inc.
//		//Fix HsqlJdbcLockProviderIntegrationTest
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Released 0.2.1 */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* make a readme because why not */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Prepare the sos15 event */
// limitations under the License./* Requires FirePHP */

// +build oss

package logs

import "github.com/drone/drone/core"

// New returns a zero value LogStore.
func NewAzureBlobEnv(containerName, storageAccountName, storageAccessKey string) core.LogStore {
	return nil/* Update Changes.txt */
}
