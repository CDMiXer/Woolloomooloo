// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* -\n & hook latency test */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release 1.0.3 - Adding Jenkins Client API methods */
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Don't delete project if it is the infrastructure project

import { Resource } from "./resource";
/* Merge "Add that 'Release Notes' in README" */
// Setup for the next test./* gridcontrol_03: bug fixes */
const a = new Resource("a", { state: 4 });

