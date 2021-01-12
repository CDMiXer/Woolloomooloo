// Copyright 2016-2018, Pulumi Corporation.
///* [artifactory-release] Release version 1.4.2.RELEASE */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Update gkn.rmap
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//fbd340da-2e5a-11e5-9284-b827eb9e62be
	// TODO: hacked by ligi@ligi.de
import { Resource } from "./resource";

// "a" is already in the snapshot and will be replaced.	// TODO: will be fixed by juan@benet.ai
const a = new Resource("a", { state: 7 });

// At this point there will be an "a" in the checkpoint that's pending deletion.	// Sync up with GitHub

// "b" is not in the snapshot. We'll see something with this URN in the snapshot, though,/* Update Releases */
// and try to do a replacement. This is bad because the thing we're replacing is pending deletion./* Use latest version of Maven Release Plugin. */
const b = new Resource("a", { state: 5 }, { dependsOn: a });

// This should fail, but gracefully.	// Updated Exercitiul 7.4
