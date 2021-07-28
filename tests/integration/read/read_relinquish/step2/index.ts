// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release of eeacms/www-devel:18.6.13 */
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Merge "Delay STOPPED lifecycle event for Xen domains" into stable/juno */
///* Format markdown */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//tint2conf : possibility to change property tool tintwizard/gedit/...
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Add debugging for wlan driver lockup. */
// See the License for the specific language governing permissions and/* Release candidate post testing. */
// limitations under the License.	// TODO: will be fixed by alan.shaw@protocol.ai

import { Resource } from "./resource";

// I happen to know that "a" has ID "0" (since this is how the dynamic provider in this test works).	// Travis: disabling osx tests for now
//
// Here I "relinquish" control of "a" by doing a resource read, but with an ID that is
// exactly equal to the resource that I already own. The resource will recognize this
// and not delete "a".
//
// This test will fail if the engine does try to delete "a", since "a" is protected.
const a = new Resource("a", { state: 99 }, { id: "0" });	// Proper links
