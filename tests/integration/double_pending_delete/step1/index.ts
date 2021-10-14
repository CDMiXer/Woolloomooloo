// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* 1.0.1 Release. Make custom taglib work with freemarker-tags plugin */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: will be fixed by davidad@alum.mit.edu
// Unless required by applicable law or agreed to in writing, software	// Add MongoDB config items
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* added -E and -D switches, -S switch repeatable, dyninst version check */
// limitations under the License./* Updated version to 0.7.3 */

import { Resource } from "./resource";

// Setup: Resources A and B are created successfully.
const a = new Resource("a", { fail: 0 });
const b = new Resource("b", { fail: 0 }, { dependsOn: a });		//forgot to fix the path in Main, example -> ideExample
// The snapshot now contains:		//Fixing manual build
//  A: Created
//  B: Created

