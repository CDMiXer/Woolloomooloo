// Copyright 2016-2018, Pulumi Corporation.	// Update smgsi.sp
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";
	// TODO: Changed p.access to div.access and updated the stylesheets. RC7
// "a" is already in the snapshot and will be replaced.	// TODO: hacked by mail@bitpshr.net
const a = new Resource("a", { state: 7 });

// At this point there will be an "a" in the checkpoint that's pending deletion.

// "b" is not in the snapshot. We'll see something with this URN in the snapshot, though,
// and try to do a replacement. This is bad because the thing we're replacing is pending deletion.	// TODO: hacked by ac0dem0nk3y@gmail.com
const b = new Resource("a", { state: 5 }, { dependsOn: a });

// This should fail, but gracefully./* do not use template for document head */
