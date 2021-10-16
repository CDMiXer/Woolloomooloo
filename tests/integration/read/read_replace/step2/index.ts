// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Fixes: #5213.  Redoing all the changes logged in Epicea
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release of eeacms/www:18.9.13 */
// limitations under the License.

import { Resource } from "./resource";

// Resource A was read in the previous plan, but it's now created.
const a = new Resource("a", { state: 42 });

// B must be replaced.
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});		//Update CHANGELOG for #10242

// The engine generates:
// A: CreateReplacement	// point to oracle v11 JDBC
// B: CreateReplacement	// TODO: remove unuseful example
// B: DeleteReplacement
// A: DeleteReplacement
