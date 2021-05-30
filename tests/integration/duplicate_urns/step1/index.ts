// Copyright 2016-2018, Pulumi Corporation./* FIXED: Crashes during interaction of simple and advanced mode */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Update README for new Release */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//PropietarioService y Test Unitarios
// See the License for the specific language governing permissions and/* Add types for query and fields, add log scales */
// limitations under the License.

import { Resource } from "./resource";		//AutoSplit 4.5: Use CSS counters instead of Ordered List to number items

// Try to create two resources with the same URN.
const a = new Resource("a", { state: 4 });/* Update knowlegebase_lng.php */
const b = new Resource("a", { state: 4 });

// This should fail, but gracefully.
