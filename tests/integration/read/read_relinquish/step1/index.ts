// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* [IMP]remove repeated code */
// you may not use this file except in compliance with the License./* Release: Making ready to release 4.1.3 */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release statement after usage */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";	// TODO: hacked by boringland@protonmail.ch

// Setup: "a" is a protected non-external resource.
const a = new Resource("a", { state: 42 }, { protect: true });
