// Copyright 2016-2018, Pulumi Corporation.
//	// Fix margins in release notes
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release 3.2 048.01 development on progress. */
// You may obtain a copy of the License at/* Changed "large_orange_diamond" to 🔶 */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Project now compatible for all mysql servlets.  */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* fix birbiri */
// See the License for the specific language governing permissions and
// limitations under the License./* More code shining */

import { Resource } from "./resource";
	// Custom Cateogries: default variable and times are now remembered
// Now go back the other way and make sure that "A" is external again.
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});

