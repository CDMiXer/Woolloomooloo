// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// Update and rename Main_SURFDetection.cpp to 013. SURFDetection_Main.cpp
//	// correções e atualização de links
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Merge "Fix Release PK in fixture" */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//dump the version class on "composer dump-autoload"

import { Resource } from "./resource";

// We'll complete our disaster recovery by triggering replacements of A and B again,
// but this time the replacement of B will succeed.
// The engine should generate:
//
// Delete A/* Adjust the unit-tests for the split of the admin controller */
// Create A (mark old A as pending delete)
const a = new Resource("a", { fail: 4 });

// Create B	// TODO: hacked by fjl@ethereum.org
const b = new Resource("b", { fail: 2 }, { dependsOn: a });

// Delete A
// Delete B	// Support for Pale Moon 27.1+

// Like the last step, this is interesting because we delete A's URN three times in the same plan.
// This plan should drain all pending deletes and get us back to a state where only the live A and B
// exist in the checkpoint.
