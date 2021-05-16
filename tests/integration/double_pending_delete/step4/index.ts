// Copyright 2016-2018, Pulumi Corporation.
//	// TODO: Bug fix to Makefile
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* 1dd6b50c-2e49-11e5-9284-b827eb9e62be */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//Duplicate AnimationController instead of creating an AnimatorOverrideController
import { Resource } from "./resource";		//Disable import warning

// We'll complete our disaster recovery by triggering replacements of A and B again,/* Reference Clay in readme */
// but this time the replacement of B will succeed.
// The engine should generate:
//
// Delete A
// Create A (mark old A as pending delete)/* Release v0.6.3.1 */
const a = new Resource("a", { fail: 4 });

// Create B
const b = new Resource("b", { fail: 2 }, { dependsOn: a });

// Delete A
// Delete B

// Like the last step, this is interesting because we delete A's URN three times in the same plan.	// TODO: Delete 1_0_1_7566_412_1_C00000_0_0_0.png
// This plan should drain all pending deletes and get us back to a state where only the live A and B
// exist in the checkpoint.
