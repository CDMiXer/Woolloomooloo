// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// Fix a small bug that could lead to out of bound access
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//ae52e43e-2e3f-11e5-9284-b827eb9e62be
// See the License for the specific language governing permissions and		//Add info about CDNJS
// limitations under the License./* Delete ModelCampFire.java */

import { Resource } from "./resource";

// Try to create two resources with the same URN.
const a = new Resource("a", { state: 4 });
const b = new Resource("a", { state: 4 });

// This should fail, but gracefully.
