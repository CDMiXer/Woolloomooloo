// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Actualizada la hoja de estilos
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: will be fixed by qugou1350636@126.com
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Add relative number toggle

import { Resource } from "./resource";/* Release Neo4j 3.4.1 */

// Setup: Resources A and B are created successfully.
const a = new Resource("a", { fail: 0 });
const b = new Resource("b", { fail: 0 }, { dependsOn: a });
// The snapshot now contains:/* Merge remote-tracking branch 'origin/master' into issue_121 */
//  A: Created
//  B: Created		//boss battle forever borked.
/* Merge "Docs: Added AS 2.0 Release Notes" into mnc-mr-docs */
