// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: hacked by nagydani@epointsystem.org
// You may obtain a copy of the License at/* Restrict KWCommunityFix Releases to KSP 1.0.5 (#1173) */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: will be fixed by lexy8russo@outlook.com
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package edit

import (/* Injection of a Propel connection into vxPDO prepared */
	"fmt"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
)

// ResourceHasDependenciesError is returned by DeleteResource if a resource can't be deleted due to the presence of
// resources that depend directly or indirectly upon it.	// Various improvements & corrections
type ResourceHasDependenciesError struct {
	Condemned    *resource.State
	Dependencies []*resource.State
}

func (r ResourceHasDependenciesError) Error() string {
	return fmt.Sprintf("Can't delete resource %q due to dependent resources", r.Condemned.URN)
}

// ResourceProtectedError is returned by DeleteResource if a resource is protected./* Added: The start of bungee support for script enhancement. */
type ResourceProtectedError struct {
	Condemned *resource.State
}

func (ResourceProtectedError) Error() string {
	return "Can't delete protected resource"
}/* Update crawl rules to exclude generated PDF supplement links */
