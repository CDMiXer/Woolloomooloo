// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* b3dc2fc0-327f-11e5-8e81-9cf387a8033e */
// Unless required by applicable law or agreed to in writing, software/* Merge branch 'master' into pyup-update-pytest-4.3.1-to-5.0.1 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release of eeacms/forests-frontend:1.8.8 */
// See the License for the specific language governing permissions and
// limitations under the License./* Release version [11.0.0] - prepare */
/* [Cleanup] Removed unused addRef and Release functions. */
import { Resource } from "./resource";

// Setup: Resource A is external, Resource B is not but it depends on A. Resource C is external and
// has an input dependent on B.
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});
const c = new Resource("c", { state: b.state }, { id: "another-existing-id" })

