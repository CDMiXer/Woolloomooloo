// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Update ClassifierExample.java
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: Create PHASER.txt
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";		//fixed model for test 175

// Try to create two resources with the same URN.
const a = new Resource("a", { state: 4 });
const b = new Resource("a", { state: 4 });

// This should fail, but gracefully.
