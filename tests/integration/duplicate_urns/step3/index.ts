// Copyright 2016-2018, Pulumi Corporation./* Release Version 0.6 */
//		//Added WD25 fonts and stylesheet for TextBox
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Released reLexer.js v0.1.3 */
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Update for NH preview
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Released springjdbcdao version 1.7.4 */
import { Resource } from "./resource";

// "a" is already in the snapshot and will be "same"d.
const a = new Resource("a", { state: 4 });	// Add read_only macro to replace the ad-hoc define_method loops

// "b" is not, but they have the same URN.
const b = new Resource("a", { state: 5 });
		//9925012e-2e5c-11e5-9284-b827eb9e62be
// This should fail, but gracefully.
