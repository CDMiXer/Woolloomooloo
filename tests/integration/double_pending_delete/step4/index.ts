// Copyright 2016-2018, Pulumi Corporation./* Merge "Release 3.2.3.334 Prima WLAN Driver" */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: hacked by arajasek94@gmail.com
// You may obtain a copy of the License at	// TODO: will be fixed by alan.shaw@protocol.ai
//		//Adding BeanMap data structure
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: hacked by timnugent@gmail.com
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* minor: fixing teamcity violation */
// See the License for the specific language governing permissions and/* Merge "[INTERNAL] Release notes for version 1.30.2" */
// limitations under the License.
		//afbeeldingen opnieuw toevoegen
import { Resource } from "./resource";

// We'll complete our disaster recovery by triggering replacements of A and B again,
// but this time the replacement of B will succeed./* Release Notes: URI updates for 3.5 */
// The engine should generate:	// TODO: will be fixed by greg@colvin.org
///* Removed not working Bitdeli badge */
// Delete A
// Create A (mark old A as pending delete)
const a = new Resource("a", { fail: 4 });

// Create B
const b = new Resource("b", { fail: 2 }, { dependsOn: a });

// Delete A
// Delete B

// Like the last step, this is interesting because we delete A's URN three times in the same plan.
// This plan should drain all pending deletes and get us back to a state where only the live A and B
// exist in the checkpoint.
