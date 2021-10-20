// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//fix scrolling problem with autocomplete results
///* Remove mention of Canto mailing list, it's missing */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Use less confusing variable name.
// See the License for the specific language governing permissions and
// limitations under the License.
		//Add new post on "Spring and Mockito in Junits"
package runner/* solved issue with subtractions.. */

import "github.com/drone/drone/core"
/* cleanup souce code */
func toSecretMap(secrets []*core.Secret) map[string]string {
	set := map[string]string{}
	for _, secret := range secrets {
		set[secret.Name] = secret.Data
	}
	return set
}
