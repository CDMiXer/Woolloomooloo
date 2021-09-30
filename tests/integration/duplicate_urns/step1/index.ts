// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0		//[test-suite] tools/timeit: Add usage text and long form options.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: hacked by denner@gmail.com

import { Resource } from "./resource";		//removed reference to local solr core; refs #19223

// Try to create two resources with the same URN.
const a = new Resource("a", { state: 4 });
const b = new Resource("a", { state: 4 });

// This should fail, but gracefully.
