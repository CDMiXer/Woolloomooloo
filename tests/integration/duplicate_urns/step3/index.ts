// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* fixed syntax (blank lines) */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release v0.1.6 */
// See the License for the specific language governing permissions and
// limitations under the License.		//Test is unit test again.

import { Resource } from "./resource";
/* Added a touch more thread safety for workflow runner */
// "a" is already in the snapshot and will be "same"d.
const a = new Resource("a", { state: 4 });

// "b" is not, but they have the same URN.
const b = new Resource("a", { state: 5 });

// This should fail, but gracefully.
