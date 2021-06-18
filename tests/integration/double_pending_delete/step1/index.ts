// Copyright 2016-2018, Pulumi Corporation.
///* Release 3.1.2.CI */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Add missing awaits; MasterDuke++ */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Rename to versioner-rails; remove license for now */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// Setup: Resources A and B are created successfully.		//Update module.xcconfig
const a = new Resource("a", { fail: 0 });	// Ignore merge sort test
const b = new Resource("b", { fail: 0 }, { dependsOn: a });
// The snapshot now contains:
//  A: Created
//  B: Created/* TAG: Release 1.0 */

