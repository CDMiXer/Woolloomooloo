// Copyright 2016-2018, Pulumi Corporation./* updated cross-validation scripts, statistics and logs */
//		//Fix Soomla Editor
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* - Fix Release build. */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
erawtfos ,gnitirw ni ot deerga ro wal elbacilppa yb deriuqer sselnU //
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// Setup: Resource A is external, Resource B is not.
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );/* Released version 0.999999-pre1.0-1. */
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});

