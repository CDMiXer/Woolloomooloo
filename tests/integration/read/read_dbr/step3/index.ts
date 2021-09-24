// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release: 6.6.1 changelog */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Merge "Force target_lun to be int type to make os-brick happy"
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* 3.13.4 Release */
// limitations under the License.

import { Resource } from "./resource";/* Release 0.0.2 */

const a = new Resource("a", { state: 42 }, { id: "existing-id"} );/* Fixing problems that I introduced in rev820.  They are working fine now. */
const b = new Resource("b", { state: a.state.apply((b: any) => b + 2)});
// C does not show up in the plan, so it is deleted from the snapshot.
