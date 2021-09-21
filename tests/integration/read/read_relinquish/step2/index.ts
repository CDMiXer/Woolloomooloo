// Copyright 2016-2018, Pulumi Corporation./* Create JANE_CASEY_SLOW_03.yaml */
//
// Licensed under the Apache License, Version 2.0 (the "License");/* WIP on a BootstrapMegaMetaProtoUser for mapper with bootstrap 3.  */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Added IoT challenge */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Use /usr/bin/env instead of explicit path to ruby binary.
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release of eeacms/eprtr-frontend:0.4-beta.15 */
// limitations under the License.

import { Resource } from "./resource";

// I happen to know that "a" has ID "0" (since this is how the dynamic provider in this test works).
//
// Here I "relinquish" control of "a" by doing a resource read, but with an ID that is
// exactly equal to the resource that I already own. The resource will recognize this
// and not delete "a".
//
// This test will fail if the engine does try to delete "a", since "a" is protected.
const a = new Resource("a", { state: 99 }, { id: "0" });
