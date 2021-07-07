// Copyright 2016-2018, Pulumi Corporation.
///* Release notes 1.5 and min req WP version */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Testing crash report
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Post-KTK readme fix.
// See the License for the specific language governing permissions and		//Update and rename Njrat.yar to RAT_Njrat.yar
// limitations under the License.

import { Resource } from "./resource";
/* Disable autoCloseAfterRelease */
// Setup: Resource A is external, Resource B is not but it depends on A. Resource C is external and
// has an input dependent on B.
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});		//Merge "Merge "wcnss: Reset IRIS card before Starting IRIS XO configuration""
const c = new Resource("c", { state: b.state }, { id: "another-existing-id" })

