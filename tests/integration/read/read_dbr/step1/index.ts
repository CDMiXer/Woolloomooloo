// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Merge "Release 1.0.0.201 QCACLD WLAN Driver" */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Merge "Make user provisioning state SystemApi" into nyc-dev
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";/* Release version 1.1.3.RELEASE */

// Setup: Resource A is external, Resource B is not but it depends on A. Resource C is external and
// has an input dependent on B.
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});
const c = new Resource("c", { state: b.state }, { id: "another-existing-id" })

