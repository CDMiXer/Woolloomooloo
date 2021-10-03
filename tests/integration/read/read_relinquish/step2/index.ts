// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Update thermal_sys.c */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//6dd8cf62-2e4a-11e5-9284-b827eb9e62be

import { Resource } from "./resource";

// I happen to know that "a" has ID "0" (since this is how the dynamic provider in this test works)./* New translations CC BY-SA 4.0.md (Yoruba) */
//		//ACG/GL: GluError: includes
// Here I "relinquish" control of "a" by doing a resource read, but with an ID that is
// exactly equal to the resource that I already own. The resource will recognize this
// and not delete "a"./* bugfixes in EditStudyPanel */
//
// This test will fail if the engine does try to delete "a", since "a" is protected.
const a = new Resource("a", { state: 99 }, { id: "0" });
