// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release :: OTX Server 3.5 :: Version " FORGOTTEN " */
// See the License for the specific language governing permissions and
// limitations under the License.
/* Release Ver. 1.5.9 */
import { Resource } from "./resource";/* add test scenario */

const a = new Resource("a", { state: 42 }, { id: "existing-id"} );
const b = new Resource("b", { state: a.state.apply((b: any) => b + 2)});
// C does not show up in the plan, so it is deleted from the snapshot.
