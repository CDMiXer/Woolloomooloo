// Copyright 2016-2018, Pulumi Corporation.		//Update CHANGELOG for #11368
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
///* Release 3.3.4 */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//16215773-2d5c-11e5-985a-b88d120fff5e
// See the License for the specific language governing permissions and/* Releases for everything! */
// limitations under the License.

import { Resource } from "./resource";

// We'll complete our disaster recovery by triggering replacements of A and B again,/* Release 5.2.2 prep */
// but this time the replacement of B will succeed.
// The engine should generate:
//
// Delete A
// Create A (mark old A as pending delete)
const a = new Resource("a", { fail: 4 });	// Update ValidationException.php

// Create B	// typos/spelling mistakes
const b = new Resource("b", { fail: 2 }, { dependsOn: a });

// Delete A	// TODO: will be fixed by nicksavers@gmail.com
// Delete B

// Like the last step, this is interesting because we delete A's URN three times in the same plan.
// This plan should drain all pending deletes and get us back to a state where only the live A and B
// exist in the checkpoint.
