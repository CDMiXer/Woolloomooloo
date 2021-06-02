// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Island generation working better, still with noticeable lag spikes */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Release any players held by a disabling plugin */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Create anti_capitalism.js
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// Setup: "a" is a protected non-external resource.	// TODO: will be fixed by arajasek94@gmail.com
const a = new Resource("a", { state: 42 }, { protect: true });/* Release v1.0.6. */
