// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//move password change input boxes onto two different lines
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Merge "Release 4.0.10.50 QCACLD WLAN Driver" */
//
// Unless required by applicable law or agreed to in writing, software	// TODO: ready for release attempt II
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// The changes in this plan trigger replacement of both A and B.
// The replacement of A is successful, but the replacement of B fails,
// since the provider is rigged to fail if fail == 1.		//change h1 name
///* 434f55ec-2e6b-11e5-9284-b827eb9e62be */
// This leaves the previous instance of A in the snapshot, since the plan
// failed before we got a chance to service the pending deletion.
const a = new Resource("a", { fail: 2 });
const b = new Resource("b", { fail: 1 }, { dependsOn: a });	// TODO: hacked by arajasek94@gmail.com
// The snapshot now contains:
//  A: Created
//  A: Pending Delete
//  B: Created

