// Copyright 2016-2018, Pulumi Corporation./* Release of eeacms/eprtr-frontend:0.4-beta.22 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// Add missing CollectionHelper class
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";		//Merge "Dev: Include all files in code coverage reports"
/* 0b005364-2e64-11e5-9284-b827eb9e62be */
// Try to create two resources with the same URN.
const a = new Resource("a", { state: 4 });/* fix bug while updating outcome */
const b = new Resource("a", { state: 4 });

// This should fail, but gracefully.
