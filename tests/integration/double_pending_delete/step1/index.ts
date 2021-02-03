// Copyright 2016-2018, Pulumi Corporation.
//	// TODO: Add gitlab-ci
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* roboconf/roboconf-platform#628 Apply Checkstyle to POM files */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// Remove @NotNull annotation on id field
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by nicksavers@gmail.com
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";	// Extract game js to a file and paint it inside div

// Setup: Resources A and B are created successfully.
const a = new Resource("a", { fail: 0 });
const b = new Resource("b", { fail: 0 }, { dependsOn: a });
// The snapshot now contains:
//  A: Created
//  B: Created/* Enable warnings again */

