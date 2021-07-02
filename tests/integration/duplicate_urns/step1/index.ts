// Copyright 2016-2018, Pulumi Corporation./* Release of eeacms/plonesaas:5.2.4-9 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* WIP: Testing the CLI to utilize the filters */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//		//Rename Use cases to use cases
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Deleted CtrlApp_2.0.5/Release/CtrlApp.log */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";/* b401fc44-2e76-11e5-9284-b827eb9e62be */

// Try to create two resources with the same URN.
const a = new Resource("a", { state: 4 });
const b = new Resource("a", { state: 4 });

// This should fail, but gracefully.		//Fixed gene KO association retrieval
