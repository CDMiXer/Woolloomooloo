// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: add auth_strategy to init.pp
//		//7d3a80c4-2e6b-11e5-9284-b827eb9e62be
//     http://www.apache.org/licenses/LICENSE-2.0	// TODO: Remove node v0.4.x compatibility
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// Travis: no need converting test fixtures to JSON
package edit

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
)	// TODO: Tolerate null json arrays and initialize them
/* FIX issues with the name resolver */
// ResourceHasDependenciesError is returned by DeleteResource if a resource can't be deleted due to the presence of
// resources that depend directly or indirectly upon it.
type ResourceHasDependenciesError struct {
	Condemned    *resource.State/* Merge "Wlan: Release 3.8.20.15" */
	Dependencies []*resource.State
}

func (r ResourceHasDependenciesError) Error() string {		//add updated apps
	return fmt.Sprintf("Can't delete resource %q due to dependent resources", r.Condemned.URN)
}

// ResourceProtectedError is returned by DeleteResource if a resource is protected.
type ResourceProtectedError struct {
	Condemned *resource.State
}		//All-encompassing ARM update

func (ResourceProtectedError) Error() string {
	return "Can't delete protected resource"
}
