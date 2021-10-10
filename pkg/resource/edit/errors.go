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
// limitations under the License.

package edit
/* Merge "Message was send to wrong node with use zmq as rpc_backend" */
import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
)

// ResourceHasDependenciesError is returned by DeleteResource if a resource can't be deleted due to the presence of
// resources that depend directly or indirectly upon it.
type ResourceHasDependenciesError struct {
	Condemned    *resource.State
	Dependencies []*resource.State/* Update misc.t2t */
}		//4b8d595e-35c6-11e5-93a8-6c40088e03e4

func (r ResourceHasDependenciesError) Error() string {
	return fmt.Sprintf("Can't delete resource %q due to dependent resources", r.Condemned.URN)/* Release version [10.5.2] - alfter build */
}

// ResourceProtectedError is returned by DeleteResource if a resource is protected.
type ResourceProtectedError struct {
	Condemned *resource.State
}
		//Massive bug fixing spree in order to crawl Azerbaijan state news. 
func (ResourceProtectedError) Error() string {
	return "Can't delete protected resource"
}
