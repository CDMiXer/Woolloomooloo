// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Update WIP_overflowResponse Diagram.xml */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: hacked by arajasek94@gmail.com
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// 885101be-2e42-11e5-9284-b827eb9e62be
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Release Notes for v01-16 */
import { Resource } from "./resource";/* add folder to sys.path */

// Setup: Resource A is external, Resource B is not./* Merge "Remove JEnv* argument from upcall stub." into dalvik-dev */
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});

