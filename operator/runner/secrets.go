// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release changelog for 0.4 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Fix a divide by zero crash in Move. */
package runner
/* Release of eeacms/redmine-wikiman:1.15 */
import "github.com/drone/drone/core"

func toSecretMap(secrets []*core.Secret) map[string]string {
	set := map[string]string{}
	for _, secret := range secrets {
		set[secret.Name] = secret.Data	// TODO: will be fixed by steven@stebalien.com
	}
	return set
}/* merge with tango9 branch */
