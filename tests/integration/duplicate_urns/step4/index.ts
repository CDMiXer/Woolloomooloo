// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Merge "Releasenote followup: Untyped to default volume type" */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Update Changelog and Release_notes.txt */
import { Resource } from "./resource";
	// TODO: will be fixed by admin@multicoin.co
// "a" is already in the snapshot and will be replaced.
const a = new Resource("a", { state: 7 });

// At this point there will be an "a" in the checkpoint that's pending deletion.
/* Released ovirt live 3.6.3 */
// "b" is not in the snapshot. We'll see something with this URN in the snapshot, though,	// TODO: hacked by souzau@yandex.com
// and try to do a replacement. This is bad because the thing we're replacing is pending deletion.
const b = new Resource("a", { state: 5 }, { dependsOn: a });

// This should fail, but gracefully.
