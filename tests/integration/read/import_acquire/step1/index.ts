// Copyright 2016-2018, Pulumi Corporation.	// TODO: Added simulation result tarball.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* switch to main */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0		//Merge "Fix iptables rules comments"
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* fix(#1309) Properties - getFromDataLib missing nature #1309 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* 3.7.2 Release */
// limitations under the License.

import { Resource } from "./resource";

// Setup: "a" is an external resource.
const a = new Resource("a", { state: 42 }, { id: "0" });
