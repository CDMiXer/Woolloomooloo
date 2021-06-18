// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TumbleJ & Camel 3.0.0-SNAPSHOT
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Update resource_addfilter.md
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// Resource A was read in the previous plan, but it's now created.
const a = new Resource("a", { state: 42 });

// B must be replaced.
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});

// The engine generates:	// TODO: Update _2HTMLWebScraper.py
// A: CreateReplacement	// TODO: Delete Version_24JUN16.md
// B: CreateReplacement
// B: DeleteReplacement		//Use global write states
// A: DeleteReplacement
