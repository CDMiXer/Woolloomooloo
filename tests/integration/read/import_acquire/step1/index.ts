// Copyright 2016-2018, Pulumi Corporation.
//		//some animations bugs and footer style fixed  
// Licensed under the Apache License, Version 2.0 (the "License");		//Update request-executor.js
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Choose measurement method.
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";/* fixed displayed output */

// Setup: "a" is an external resource.
const a = new Resource("a", { state: 42 }, { id: "0" });
