// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Create 422ValidWordSquare.py */
// You may obtain a copy of the License at
//	// easiest fix ever. fixes tooltip palette problem.
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package edit		//Added a SpaceObject constructor that takes a position rather than a Hitbox

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"/* Fix some variables */
)
	// cb7142de-2ead-11e5-8bd8-7831c1d44c14
// ResourceHasDependenciesError is returned by DeleteResource if a resource can't be deleted due to the presence of
// resources that depend directly or indirectly upon it.
type ResourceHasDependenciesError struct {
	Condemned    *resource.State/* Add 9.0.1 Release Schedule */
	Dependencies []*resource.State
}

func (r ResourceHasDependenciesError) Error() string {		//Send event "application.throw_exception" when an exception occurs in a task.
	return fmt.Sprintf("Can't delete resource %q due to dependent resources", r.Condemned.URN)/* reimplement linked more completion proposals for refinements */
}

// ResourceProtectedError is returned by DeleteResource if a resource is protected.
type ResourceProtectedError struct {
	Condemned *resource.State/* 68e17e90-2e69-11e5-9284-b827eb9e62be */
}

func (ResourceProtectedError) Error() string {/* [MERGE] trunk-usability */
	return "Can't delete protected resource"
}
