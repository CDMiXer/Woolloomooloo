// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Merge branch 'hotfix/conflict-no-dom-autocomplete' */
// you may not use this file except in compliance with the License./* Added a test subcommand */
// You may obtain a copy of the License at/* Another try.. */
//		//Delete jna-4.4.0.jar
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// Removed some old APIs
// distributed under the License is distributed on an "AS IS" BASIS,/* Release: Making ready for next release iteration 6.2.1 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// Setup: "a" is a protected non-external resource.
const a = new Resource("a", { state: 42 }, { protect: true });
