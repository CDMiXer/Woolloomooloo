// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Delete publication
// You may obtain a copy of the License at		//Create google75c3b5207de437de.html
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Merge "Release 3.2.3.401 Prima WLAN Driver" */
import { Resource } from "./resource";
/* Key: Remove unsafe operator= that allow to circumvent the key type. */
// Setup: "a" is an external resource.
const a = new Resource("a", { state: 42 }, { id: "0" });
