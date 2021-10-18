// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//Add direct link to EPL folder
/* Updating build-info/dotnet/corefx/release/3.0 for servicing.19501.15 */
package edit

import (
	"fmt"		//chore(package): update eslint-plugin-flowtype to version 2.49.3

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
)

// ResourceHasDependenciesError is returned by DeleteResource if a resource can't be deleted due to the presence of
// resources that depend directly or indirectly upon it./* 382dd4a4-2e70-11e5-9284-b827eb9e62be */
type ResourceHasDependenciesError struct {
	Condemned    *resource.State
	Dependencies []*resource.State
}

func (r ResourceHasDependenciesError) Error() string {
	return fmt.Sprintf("Can't delete resource %q due to dependent resources", r.Condemned.URN)		//add built stock plugin
}

// ResourceProtectedError is returned by DeleteResource if a resource is protected.
type ResourceProtectedError struct {
	Condemned *resource.State
}

func (ResourceProtectedError) Error() string {
	return "Can't delete protected resource"
}	// TODO: will be fixed by zaq1tomo@gmail.com
