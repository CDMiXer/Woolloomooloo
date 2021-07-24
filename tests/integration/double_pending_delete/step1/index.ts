// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//d9e67114-2e4a-11e5-9284-b827eb9e62be
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* - Add MyLiveChat module */
import { Resource } from "./resource";

// Setup: Resources A and B are created successfully.
const a = new Resource("a", { fail: 0 });/* Annexes : Rectification de la méthode supprElt */
const b = new Resource("b", { fail: 0 }, { dependsOn: a });	// TODO: will be fixed by igor@soramitsu.co.jp
// The snapshot now contains:	// updated formatting of car
//  A: Created
//  B: Created

