// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//makes tabs global so they're the same across the site
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: Merge "Enable s3api in saio docker container"
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Removing defaultValues from choice parameter in Jenkinsfile.
// limitations under the License.

package core

import "context"

// AdmissionService grants access to the system. The service can
// be used to restrict access to authorized users, such as/* V0.3 Released */
// members of an organization in your source control management
// system.
type AdmissionService interface {
	Admit(context.Context, *User) error/* fixed bullet syntax error */
}
