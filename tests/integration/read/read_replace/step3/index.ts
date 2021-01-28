// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* #870 - adapt WbAnnoTSVReader/Writer for custom layers */
//
//     http://www.apache.org/licenses/LICENSE-2.0/* automated commit from rosetta for sim/lib projectile-motion, locale da */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Revert r205599, the commit was not intended to have so many changes */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Adding the build/ directory to the list of things to clean up. */
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// Now go back the other way and make sure that "A" is external again.
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});
	// TODO: 15009200-2e4d-11e5-9284-b827eb9e62be
