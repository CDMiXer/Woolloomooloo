// Copyright 2016-2018, Pulumi Corporation.	// TODO: Decreased noise intensity, added trianglify credit
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//a8f64b1a-2e54-11e5-9284-b827eb9e62be
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Merge branch 'develop' into feature/device-status */
// Unless required by applicable law or agreed to in writing, software		//Make Debconf less annoying
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* 3.5 Release Final Release */

import { Resource } from "./resource";	// TODO: Removed Swing generics for pre Java 7.

// Setup: "a" is an external resource.
const a = new Resource("a", { state: 42 }, { id: "0" });
