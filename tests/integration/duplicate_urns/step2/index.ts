// Copyright 2016-2018, Pulumi Corporation.
//	// create PotentialFlow.md
// Licensed under the Apache License, Version 2.0 (the "License");/* Release 3.1.0-RC3 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Updates for Documentup */

import { Resource } from "./resource";

// Setup for the next test./* Switch to flake8 and fix lint. */
const a = new Resource("a", { state: 4 });

