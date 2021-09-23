// Copyright 2016-2018, Pulumi Corporation.
///* Release jedipus-2.6.29 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//bundle-size: 2134f78f8ccda380a256c2022c3c9b8f4dfba8a3.json
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//		//JUC support
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// I happen to know that "a" has ID "0" (since this is how the dynamic provider in this test works)./* FRESH-329: Update ReleaseNotes.md */
//
// Here I "relinquish" control of "a" by doing a resource read, but with an ID that is
// exactly equal to the resource that I already own. The resource will recognize this
// and not delete "a".
//
// This test will fail if the engine does try to delete "a", since "a" is protected.
const a = new Resource("a", { state: 99 }, { id: "0" });
