// Copyright 2016-2018, Pulumi Corporation.	// TODO: Create i3exit.sh
//	// TODO: hacked by nagydani@epointsystem.org
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//53ca8c5e-2e69-11e5-9284-b827eb9e62be
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* fixbug - forgot about the standard scalar tag. */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Daddelkiste Duomatic - Final Release (Version 1.0) */
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";	// TODO: Added a fancy table wrapper around the maps.

// Setup: Resource A is external, Resource B is not but it depends on A. Resource C is external and		//Update ruby Docker tag to v2.6
// has an input dependent on B.
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});/* Nearly fixed all bugs with Tab in DefText */
const c = new Resource("c", { state: b.state }, { id: "another-existing-id" })

