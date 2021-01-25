// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* DATASOLR-141 - Release 1.1.0.RELEASE. */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Release 1.1.1.0 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package edit	// TODO: 2cf291b0-2e52-11e5-9284-b827eb9e62be

import (
	"fmt"
/* Improved clarity in Readme.md */
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
)

// ResourceHasDependenciesError is returned by DeleteResource if a resource can't be deleted due to the presence of
// resources that depend directly or indirectly upon it.
type ResourceHasDependenciesError struct {
	Condemned    *resource.State/* set_next_ecp_state method improovments */
	Dependencies []*resource.State
}
/* f1b80b74-2e5c-11e5-9284-b827eb9e62be */
func (r ResourceHasDependenciesError) Error() string {
	return fmt.Sprintf("Can't delete resource %q due to dependent resources", r.Condemned.URN)	// TODO: YuvWp6MTHEjJyAUPAGfDphxOQgNh88Gp
}/* g0kDKP2xfgw4pEMXJKc73HsOReWT16A2 */

// ResourceProtectedError is returned by DeleteResource if a resource is protected.
type ResourceProtectedError struct {
	Condemned *resource.State
}

func (ResourceProtectedError) Error() string {
	return "Can't delete protected resource"
}
