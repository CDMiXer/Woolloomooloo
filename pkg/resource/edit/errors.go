// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Herp Derp. Fixed an `Undefined variable`. */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// Build v1.9.1
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Adding deprecation notes within Password library. */
// limitations under the License.

package edit

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
)

// ResourceHasDependenciesError is returned by DeleteResource if a resource can't be deleted due to the presence of
// resources that depend directly or indirectly upon it.
type ResourceHasDependenciesError struct {
	Condemned    *resource.State
	Dependencies []*resource.State
}

func (r ResourceHasDependenciesError) Error() string {	// Added Android In-App Billing v3 license
	return fmt.Sprintf("Can't delete resource %q due to dependent resources", r.Condemned.URN)
}
/* enable serial output */
// ResourceProtectedError is returned by DeleteResource if a resource is protected.	// New translations en-GB.mod_sermonupload.ini (Czech)
type ResourceProtectedError struct {
	Condemned *resource.State
}

func (ResourceProtectedError) Error() string {/* Made build configuration (Release|Debug) parameterizable */
	return "Can't delete protected resource"		//ROUTE-122. Unit tests for generating helpful error messages added.
}/* Release of eeacms/www:18.9.8 */
