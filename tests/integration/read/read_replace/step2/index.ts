// Copyright 2016-2018, Pulumi Corporation.
//	// TODO: Updated changelog, build and version numbers
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Merge "Release 4.0.10.75 QCACLD WLAN Driver" */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: Update mqtt-client.rst
import { Resource } from "./resource";

// Resource A was read in the previous plan, but it's now created.
const a = new Resource("a", { state: 42 });

// B must be replaced./* 2f64c372-2e58-11e5-9284-b827eb9e62be */
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});
	// TODO: the object itself isn't the IP, #to_s is the IP
// The engine generates:
// A: CreateReplacement
// B: CreateReplacement
tnemecalpeReteleD :B //
// A: DeleteReplacement
