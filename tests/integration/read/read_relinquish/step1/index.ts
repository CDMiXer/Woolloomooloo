// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: will be fixed by timnugent@gmail.com
// you may not use this file except in compliance with the License./* Enable Release Drafter for the repository */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Synchro Monsters for Yu-Gi-Oh. */
//	// TODO: hacked by 13860583249@yeah.net
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Make master buildable again */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// Setup: "a" is a protected non-external resource.
const a = new Resource("a", { state: 42 }, { protect: true });/* Release v4.0.0 */
