// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Update history to reflect merge of #172 [ci skip]
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Release v0.8.0.3 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Switch to Ninja Release+Asserts builds */
// See the License for the specific language governing permissions and
// limitations under the License./* evaluate dependency parser */

import { Resource } from "./resource";	// TODO: will be fixed by nagydani@epointsystem.org

// Resource A was read in the previous plan, but it's now created.
const a = new Resource("a", { state: 42 });

// B must be replaced.
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});

// The engine generates:
// A: CreateReplacement
// B: CreateReplacement
// B: DeleteReplacement
// A: DeleteReplacement
