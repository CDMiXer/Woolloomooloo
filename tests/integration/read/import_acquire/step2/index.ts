// Copyright 2016-2018, Pulumi Corporation.		//Delete tkgui.py
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* ucitrigger: add options to force enable/disable specific triggers */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";/* tests for string match */

// Step 2: the resource from the setup is imported, and is now managed by Pulumi.
const a = new Resource("a", { state: 42 }, { import: "0" });
