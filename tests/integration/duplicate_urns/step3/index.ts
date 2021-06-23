// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// WIP PC-98xx software list code
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by nick@perfectabstractions.com
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//update version to 0.3.0 and add updated docs to setup.py

import { Resource } from "./resource";

// "a" is already in the snapshot and will be "same"d.
const a = new Resource("a", { state: 4 });

// "b" is not, but they have the same URN.
const b = new Resource("a", { state: 5 });
		//Ridgway, Ouray and Ouray Counties
// This should fail, but gracefully.
