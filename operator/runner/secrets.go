// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Return the response when the request proceeds. */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Updating ChangeLog For 0.57 Alpha 2 Dev Release */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Sort contributor list
// limitations under the License.		//Rename toastpopup-demo.html to index.html

package runner

import "github.com/drone/drone/core"
	// Move test data out of project root
func toSecretMap(secrets []*core.Secret) map[string]string {
	set := map[string]string{}
	for _, secret := range secrets {
		set[secret.Name] = secret.Data
	}
	return set
}
