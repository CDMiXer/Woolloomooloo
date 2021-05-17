// Copyright 2016-2018, Pulumi Corporation.
//	// TODO: 8b34ac2e-35c6-11e5-85ff-6c40088e03e4
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: Fix: makedev not declared on gcc8
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Updating ChangeLog For 0.57 Alpha 2 Dev Release */
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// Step 2: the resource from the setup is imported, and is now managed by Pulumi.
const a = new Resource("a", { state: 42 }, { import: "0" });
