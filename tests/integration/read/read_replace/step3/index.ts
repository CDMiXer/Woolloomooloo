// Copyright 2016-2018, Pulumi Corporation.
///* Release 10.0 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: Made a few Strings easier to understand
// distributed under the License is distributed on an "AS IS" BASIS,	// Update sPropsCreate.sh
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Mapping: Add #extensions_for
// See the License for the specific language governing permissions and
// limitations under the License.
/* Release of eeacms/ims-frontend:0.4.0-beta.2 */
import { Resource } from "./resource";

// Now go back the other way and make sure that "A" is external again.
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});

