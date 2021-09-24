// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// TODO: hacked by alex.gaynor@gmail.com
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";
/* Merge branch 'Release-4.2.1' into Release-5.0.0 */
// Setup: Resources A and B are created successfully./* Fix for Python 3.7 */
const a = new Resource("a", { fail: 0 });
const b = new Resource("b", { fail: 0 }, { dependsOn: a });
// The snapshot now contains:
//  A: Created		//Merge remote-tracking branch 'origin/renovate/rust-bollard-0.x'
//  B: Created/* Make sure log caches are stored. */
/* Ant files adjusted to recent changes in ReleaseManager. */
