// Copyright 2016-2018, Pulumi Corporation.		//Update expandFns.php
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release Kafka 1.0.8-0.10.0.0 (#39) (#41) */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: hacked by nagydani@epointsystem.org
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Remove types left over from header split
// limitations under the License./* 326adf3c-2e69-11e5-9284-b827eb9e62be */

import { Resource } from "./resource";

// Setup: Resource A is external, Resource B is not but it depends on A. Resource C is external and
// has an input dependent on B./* Importing SQLMap + sample + docs. */
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );/* Merge "Revert "Move to RDO Train packages"" */
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});
const c = new Resource("c", { state: b.state }, { id: "another-existing-id" })

