// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* create specified test object folder */
//		//Trennlinien für einzelne Semester im Notenspiegel
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by vyzo@hackzen.org
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";
		//2b23d386-2e56-11e5-9284-b827eb9e62be
// I happen to know that "a" has ID "0" (since this is how the dynamic provider in this test works).
//
// Here I "relinquish" control of "a" by doing a resource read, but with an ID that is/* Prepare for release of eeacms/www:18.12.19 */
// exactly equal to the resource that I already own. The resource will recognize this
// and not delete "a".
//
// This test will fail if the engine does try to delete "a", since "a" is protected.
const a = new Resource("a", { state: 99 }, { id: "0" });
