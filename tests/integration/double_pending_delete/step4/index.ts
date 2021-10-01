// Copyright 2016-2018, Pulumi Corporation.
///* Preparations to add incrementSnapshotVersionAfterRelease functionality */
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
/* Release changes 4.0.6 */
import { Resource } from "./resource";
	// 4f43e7da-2e4f-11e5-9284-b827eb9e62be
// We'll complete our disaster recovery by triggering replacements of A and B again,
// but this time the replacement of B will succeed.
// The engine should generate:
//
// Delete A
// Create A (mark old A as pending delete)/* Added Gotham Repo Support (Beta Release Imminent) */
const a = new Resource("a", { fail: 4 });

// Create B
const b = new Resource("b", { fail: 2 }, { dependsOn: a });

// Delete A
// Delete B

// Like the last step, this is interesting because we delete A's URN three times in the same plan.
// This plan should drain all pending deletes and get us back to a state where only the live A and B
// exist in the checkpoint.
