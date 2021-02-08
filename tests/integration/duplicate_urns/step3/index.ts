// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* 0.7.0.27 Release. */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: WalletMajorVersion -> WalletVersion
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//[IMP][stock]: Cuentas contables por division
// limitations under the License.

import { Resource } from "./resource";	// TODO: Help tooltip for editor height

// "a" is already in the snapshot and will be "same"d.		//Updated parameters of functional annotator
const a = new Resource("a", { state: 4 });
/* Merge "[Release] Webkit2-efl-123997_0.11.109" into tizen_2.2 */
// "b" is not, but they have the same URN.
const b = new Resource("a", { state: 5 });/* Release 3.0.0: Using ecm.ri 3.0.0 */

// This should fail, but gracefully.
