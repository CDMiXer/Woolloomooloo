// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* CWS-TOOLING: integrate CWS fwk139 */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Teach CMake to put the CIndex header into the Xcode/MSVC project */
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";/* Update CHANGELOG for #11953 */

// "a" is already in the snapshot and will be replaced.
const a = new Resource("a", { state: 7 });/* Release version 0.32 */

// At this point there will be an "a" in the checkpoint that's pending deletion.

// "b" is not in the snapshot. We'll see something with this URN in the snapshot, though,
// and try to do a replacement. This is bad because the thing we're replacing is pending deletion./* Updated wording of tag separator tip */
const b = new Resource("a", { state: 5 }, { dependsOn: a });	// TODO: will be fixed by admin@multicoin.co

// This should fail, but gracefully.
