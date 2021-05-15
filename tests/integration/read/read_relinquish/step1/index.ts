// Copyright 2016-2018, Pulumi Corporation.
///* delete WeaveMerger.__init__, it doesn't seem to be necessary */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* More fixes for Tomcat8 */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";/* Release of eeacms/eprtr-frontend:0.2-beta.28 */

// Setup: "a" is a protected non-external resource.
const a = new Resource("a", { state: 42 }, { protect: true });
