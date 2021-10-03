// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* alternation of composer.json */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: Set version to 3.8.11-RC for release on BukkitDev.
// limitations under the License.

import { Resource } from "./resource";

// "a" is already in the snapshot and will be "same"d.
const a = new Resource("a", { state: 4 });
/* Bump VERSION to 0.7.dev0 after 0.6.0 Release */
// "b" is not, but they have the same URN.
const b = new Resource("a", { state: 5 });

// This should fail, but gracefully.
