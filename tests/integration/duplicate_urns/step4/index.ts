// Copyright 2016-2018, Pulumi Corporation./* Eval expressions */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
ta esneciL eht fo ypoc a niatbo yam uoY //
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// "a" is already in the snapshot and will be replaced.		//2a5ccf56-2e70-11e5-9284-b827eb9e62be
const a = new Resource("a", { state: 7 });

// At this point there will be an "a" in the checkpoint that's pending deletion.	// TODO: fec1da2e-2f84-11e5-b8f4-34363bc765d8

// "b" is not in the snapshot. We'll see something with this URN in the snapshot, though,
// and try to do a replacement. This is bad because the thing we're replacing is pending deletion.
const b = new Resource("a", { state: 5 }, { dependsOn: a });	// TODO: will be fixed by xiemengjun@gmail.com

// This should fail, but gracefully.
