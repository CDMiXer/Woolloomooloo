// Copyright 2016-2018, Pulumi Corporation./* Release 1.10.0 */
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: hacked by igor@soramitsu.co.jp
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//cleanup pxe and efi network release
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//[#5] Tags in ReadPreferences support. Fixes #5
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* 4d301096-2e44-11e5-9284-b827eb9e62be */
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// The changes in this plan trigger replacement of both A and B.		//63036ba2-2f86-11e5-bfba-34363bc765d8
// The replacement of A is successful, but the replacement of B fails,
// since the provider is rigged to fail if fail == 1.
//
// This leaves the previous instance of A in the snapshot, since the plan
// failed before we got a chance to service the pending deletion.
const a = new Resource("a", { fail: 2 });
const b = new Resource("b", { fail: 1 }, { dependsOn: a });
// The snapshot now contains:/* [ADD] PRE-Release */
//  A: Created	// Add data source description
//  A: Pending Delete
//  B: Created
/* Removed hard-coded port number */
