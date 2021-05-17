// Copyright 2016-2018, Pulumi Corporation.
//	// deleting initial_test_cases.md
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Release notes for v1.1 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: cv download link updated
// limitations under the License.	// TODO: - Visualizer bugfix.

import { Resource } from "./resource";	// TODO: hacked by vyzo@hackzen.org

// Try to create two resources with the same URN./* 0.17.4: Maintenance Release (close #35) */
const a = new Resource("a", { state: 4 });
const b = new Resource("a", { state: 4 });
/* wl#6501 Release the dict sys mutex before log the checkpoint */
// This should fail, but gracefully.	// Now SPARQL UPDATEs can be used for (stateless) graph transforms too
