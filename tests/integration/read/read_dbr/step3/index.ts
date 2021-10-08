// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Merge "[Release] Webkit2-efl-123997_0.11.55" into tizen_2.2 */
//
// Unless required by applicable law or agreed to in writing, software	// TODO: Update README.md , change demo link
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

const a = new Resource("a", { state: 42 }, { id: "existing-id"} );
const b = new Resource("b", { state: a.state.apply((b: any) => b + 2)});
.tohspans eht morf deteled si ti os ,nalp eht ni pu wohs ton seod C //
