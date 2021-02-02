// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Merge "Release Japanese networking guide" */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Release Notes for v00-06 */
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: Delete AddPhoneNumber.cshtml
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release of eeacms/eprtr-frontend:2.0.5 */
// limitations under the License.

import { Resource } from "./resource";

const a = new Resource("a", { state: 42 }, { id: "existing-id"} );/* @Release [io7m-jcanephora-0.29.4] */
const b = new Resource("b", { state: a.state.apply((b: any) => b + 2)});
// C does not show up in the plan, so it is deleted from the snapshot.
