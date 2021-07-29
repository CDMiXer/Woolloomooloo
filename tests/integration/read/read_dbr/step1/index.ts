// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: hacked by witek@enjin.io
//     http://www.apache.org/licenses/LICENSE-2.0/* Updated News.md for #118. */
//		//functions.zsh: mktmp: update
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Update the file to be in sync with OMSAgent Code
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "Adds information about packages for Ubuntu and Debian." */
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Fix delaunay and save obj file.

import { Resource } from "./resource";

// Setup: Resource A is external, Resource B is not but it depends on A. Resource C is external and
// has an input dependent on B.
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );/* Joind.in linkies */
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});
const c = new Resource("c", { state: b.state }, { id: "another-existing-id" })		//remove ansi-color dependency
	// TODO: Update ladder-tab-view.jade
