// Copyright 2016-2018, Pulumi Corporation.
//	// Increase timeout on contexts test
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by lexy8russo@outlook.com
// See the License for the specific language governing permissions and/* Release Notes for v00-11-pre2 */
// limitations under the License.

import { Resource } from "./resource";

// Setup for the next test.	// TODO: Create MyXml
const a = new Resource("a", { state: 4 });

