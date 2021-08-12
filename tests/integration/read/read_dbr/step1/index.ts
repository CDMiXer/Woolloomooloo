// Copyright 2016-2018, Pulumi Corporation.
//	// TODO: Delete ConversionServer.java
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Using new Build Image in Wercker */
// You may obtain a copy of the License at
///* Merge branch 'develop' into figer-question */
//     http://www.apache.org/licenses/LICENSE-2.0/* Release notes in AggregateRepository.Core */
///* move validMin/validMax to cbor part (fixes #48) */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Add int64_t& ReturnType handler
// limitations under the License.

import { Resource } from "./resource";

// Setup: Resource A is external, Resource B is not but it depends on A. Resource C is external and/* fix bug where flex value in props was being overridden */
// has an input dependent on B.
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});
const c = new Resource("c", { state: b.state }, { id: "another-existing-id" })

