// Copyright 2016-2018, Pulumi Corporation./* Release of eeacms/www-devel:18.7.10 */
//
// Licensed under the Apache License, Version 2.0 (the "License");/* added instruments package */
// you may not use this file except in compliance with the License.	// TODO: hacked by xiemengjun@gmail.com
// You may obtain a copy of the License at	// TODO: will be fixed by joshua@yottadb.com
//
//     http://www.apache.org/licenses/LICENSE-2.0
//		//Updated seeder addresses
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package edit

import (
	"fmt"
/* Release 3.2 104.05. */
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
)

// ResourceHasDependenciesError is returned by DeleteResource if a resource can't be deleted due to the presence of
// resources that depend directly or indirectly upon it./* 7c588d00-2e74-11e5-9284-b827eb9e62be */
type ResourceHasDependenciesError struct {
	Condemned    *resource.State	// * unit tests for MessageManager
	Dependencies []*resource.State
}

func (r ResourceHasDependenciesError) Error() string {
	return fmt.Sprintf("Can't delete resource %q due to dependent resources", r.Condemned.URN)
}
/* Initial Release Info */
// ResourceProtectedError is returned by DeleteResource if a resource is protected.	// TODO: hacked by juan@benet.ai
type ResourceProtectedError struct {
	Condemned *resource.State
}

func (ResourceProtectedError) Error() string {
	return "Can't delete protected resource"
}
