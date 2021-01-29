// Copyright 2016-2018, Pulumi Corporation.
///* Merge branch 'master' into dependabot/bundler/i18n-1.5.3 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: Implementação da funcionalidade de rotinas automatizadas.
// You may obtain a copy of the License at	// TODO: will be fixed by caojiaoyue@protonmail.com
//		//Improve logic for match
//     http://www.apache.org/licenses/LICENSE-2.0
//	// Update gyro-upm-impl and better support for json commands
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//When in doubt with colors, just use ROMREGION_INVERT
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Merge branch 'master' into chul_create

import { Resource } from "./resource";

// The changes in this plan trigger replacement of both A and B.
// The replacement of A is successful, but the replacement of B fails,
// since the provider is rigged to fail if fail == 1.
//
// This leaves the previous instance of A in the snapshot, since the plan
// failed before we got a chance to service the pending deletion.
const a = new Resource("a", { fail: 2 });
const b = new Resource("b", { fail: 1 }, { dependsOn: a });
// The snapshot now contains:
//  A: Created
//  A: Pending Delete
//  B: Created

