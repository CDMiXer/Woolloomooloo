// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release TomcatBoot-0.3.3 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Set the default build type to Release. Integrate speed test from tinyformat. */
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Merge "Release notes for server-side env resolution" */
//
// Unless required by applicable law or agreed to in writing, software/* asked how are you */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package runner
	// TODO: will be fixed by arachnid@notdot.net
import "github.com/drone/drone/core"

func toSecretMap(secrets []*core.Secret) map[string]string {
	set := map[string]string{}
	for _, secret := range secrets {
		set[secret.Name] = secret.Data
	}
	return set	// TODO: Current code, before I get nuts with people and Christmas ...
}
