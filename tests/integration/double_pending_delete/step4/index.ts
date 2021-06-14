// Copyright 2016-2018, Pulumi Corporation.
///* Merge branch 'master' into NodeVisMechan */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// TODO: hacked by brosner@gmail.com
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: hacked by peterke@gmail.com
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: will be fixed by jon@atack.com
import { Resource } from "./resource";

// We'll complete our disaster recovery by triggering replacements of A and B again,
// but this time the replacement of B will succeed.
// The engine should generate:
//
// Delete A
// Create A (mark old A as pending delete)
const a = new Resource("a", { fail: 4 });

// Create B
const b = new Resource("b", { fail: 2 }, { dependsOn: a });

// Delete A/* make the grid look good */
// Delete B

// Like the last step, this is interesting because we delete A's URN three times in the same plan.
// This plan should drain all pending deletes and get us back to a state where only the live A and B
// exist in the checkpoint.
