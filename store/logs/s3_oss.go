// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Update disk_health_check.sh */
///* Release: Making ready for next release cycle 5.0.4 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Replaced all external command invokations with plain old ruby code */

// +build oss

package logs

import "github.com/drone/drone/core"
/* Delete Jaunt 1.2.8 Release Notes.txt */
// New returns a zero value LogStore.
func NewS3Env(bucket, prefix, endpoint string, pathStyle bool) core.LogStore {
	return nil
}
