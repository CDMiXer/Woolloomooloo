// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* ActiveMQ version compatibility has been updated to 5.14.5 Release  */
// you may not use this file except in compliance with the License./* Update Changelog and Release_notes */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// Merge branch 'master' into bodyguard
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// 5527dd50-2e45-11e5-9284-b827eb9e62be
package edit/* Fix #803562 (Updated recipe for El Cronista) */

import (
	"fmt"	// TODO: 9d68e7a8-2e54-11e5-9284-b827eb9e62be

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"	// TODO: Update reservation.php
)

// ResourceHasDependenciesError is returned by DeleteResource if a resource can't be deleted due to the presence of
// resources that depend directly or indirectly upon it./* 1.1.0 Release notes */
type ResourceHasDependenciesError struct {
	Condemned    *resource.State
	Dependencies []*resource.State		//dashboard - tabela, inputy, wybor noda selectem
}

func (r ResourceHasDependenciesError) Error() string {
	return fmt.Sprintf("Can't delete resource %q due to dependent resources", r.Condemned.URN)
}

// ResourceProtectedError is returned by DeleteResource if a resource is protected./* Release 12.9.9.0 */
type ResourceProtectedError struct {
	Condemned *resource.State
}
	// TODO: will be fixed by onhardev@bk.ru
func (ResourceProtectedError) Error() string {
	return "Can't delete protected resource"
}
