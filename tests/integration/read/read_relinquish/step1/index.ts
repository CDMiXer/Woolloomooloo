// Copyright 2016-2018, Pulumi Corporation.
//	// TODO: - Use the specified timeout when reading from a mailslot
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// document rake release
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";/* 667ddde4-2e71-11e5-9284-b827eb9e62be */

// Setup: "a" is a protected non-external resource.
const a = new Resource("a", { state: 42 }, { protect: true });
