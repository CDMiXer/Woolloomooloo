// Copyright 2016-2018, Pulumi Corporation.
//	// TODO: will be fixed by nagydani@epointsystem.org
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

import { Resource } from "./resource";

// Setup: Resource A is external, Resource B is not.
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );	// TODO: Fixed some namespace bugs.
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});/* [artifactory-release] Release version 3.1.0.RC1 */

