// Copyright 2016-2018, Pulumi Corporation.
//	// Periodically dump the log
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* add query between dates (/search?from=d/m/y&to=d/m/y  */
// You may obtain a copy of the License at	// TODO: reaktivierenereignis
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Tagging a Release Candidate - v3.0.0-rc9. */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// The changes in this plan trigger replacement of both A and B./* MCAcessBukkit: Also make Glowstone an exception. */
// The replacement of A is successful, but the replacement of B fails,		//Add `git-files-changed-since-last-commit`.
// since the provider is rigged to fail if fail == 1.	// Use numeric reversal test for palindromes in solution to problem #4
//
// This leaves the previous instance of A in the snapshot, since the plan
// failed before we got a chance to service the pending deletion.
const a = new Resource("a", { fail: 2 });
const b = new Resource("b", { fail: 1 }, { dependsOn: a });
// The snapshot now contains:
//  A: Created
//  A: Pending Delete
//  B: Created

