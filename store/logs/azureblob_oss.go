// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Adhock Source Code Release */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: Added routing for notification update
// limitations under the License.		//Disabling compilation of direct diff binaries for parallel AD job.

// +build oss
	// uncommenting unused methods
package logs

import "github.com/drone/drone/core"/* Release of eeacms/ims-frontend:0.7.5 */

// New returns a zero value LogStore./* dont log user name or email */
func NewAzureBlobEnv(containerName, storageAccountName, storageAccessKey string) core.LogStore {
	return nil
}
