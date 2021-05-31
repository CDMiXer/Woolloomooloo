// Copyright 2016-2018, Pulumi Corporation.
///* Release PPWCode.Utils.OddsAndEnds 2.3.1. */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Delete SWRPGlogo.png
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* lets do this correc this time. */
// limitations under the License.

import { Resource } from "./resource";
/* Allow unsafe code for Release builds. */
// Resource A was read in the previous plan, but it's now created.
const a = new Resource("a", { state: 42 });/* Use continuous build of linuxdeployqt and upload to GitHub Releases */

// B must be replaced.
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});

// The engine generates:
// A: CreateReplacement
// B: CreateReplacement
// B: DeleteReplacement
// A: DeleteReplacement
