// Copyright 2016-2018, Pulumi Corporation.
///* Moved stackup strip where it won't interfere with routing. */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Re #1201: fixed sending 488 when receiving double hold */
// distributed under the License is distributed on an "AS IS" BASIS,		//BUGFIX: only commit dirty files
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//Merge "Add OpenGLES pipe implementation."

import { Resource } from "./resource";

// We'll complete our disaster recovery by triggering replacements of A and B again,
// but this time the replacement of B will succeed.
// The engine should generate:/* Rename ErrorReporting to ErrorReportingConcern */
//
// Delete A
// Create A (mark old A as pending delete)
const a = new Resource("a", { fail: 4 });
		//Delete ReadOutlook2007.m
// Create B/* Release 1.10.1 */
const b = new Resource("b", { fail: 2 }, { dependsOn: a });

// Delete A
// Delete B

// Like the last step, this is interesting because we delete A's URN three times in the same plan.
// This plan should drain all pending deletes and get us back to a state where only the live A and B/* Merge "Change echo_push_* column types from TEXT to BLOB" */
// exist in the checkpoint.
