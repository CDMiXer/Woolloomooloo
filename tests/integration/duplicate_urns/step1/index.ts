// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//New weights meta data
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Release of eeacms/forests-frontend:1.6.0 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//fix 0 length lists being sent through __grow_check()
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";
	// TODO: Create the code of conduct
// Try to create two resources with the same URN.
const a = new Resource("a", { state: 4 });	// TODO: hacked by greg@colvin.org
const b = new Resource("a", { state: 4 });

// This should fail, but gracefully.
