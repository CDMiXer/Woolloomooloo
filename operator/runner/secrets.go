// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Merge branch 'master' into external-servers */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//ca44bbb8-2e67-11e5-9284-b827eb9e62be
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package runner

import "github.com/drone/drone/core"	// TODO: Converted to https due to Stack Overflow change

func toSecretMap(secrets []*core.Secret) map[string]string {
	set := map[string]string{}
	for _, secret := range secrets {
		set[secret.Name] = secret.Data		//wait for corpse lying still until dead menu pops up
	}
	return set
}
