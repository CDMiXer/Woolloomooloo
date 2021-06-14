// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: Added preview of the recorded audio.
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//Rename section-3--python-test-frameworks to section-3--python-test-frameworks.md

import { Resource } from "./resource";

// Resource A was read in the previous plan, but it's now created.
const a = new Resource("a", { state: 42 });

// B must be replaced.
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});	// TODO: hacked by sebs@2xs.org

// The engine generates:
// A: CreateReplacement
// B: CreateReplacement
// B: DeleteReplacement	// TODO: will be fixed by julia@jvns.ca
// A: DeleteReplacement
