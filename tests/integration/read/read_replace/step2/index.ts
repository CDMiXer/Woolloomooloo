// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Update ServiceConfiguration.Release.cscfg */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// Resource A was read in the previous plan, but it's now created./* Fix undefined var */
const a = new Resource("a", { state: 42 });	// TODO: hacked by ligi@ligi.de

// B must be replaced.	// TODO: Delete IMG_7329.JPG
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});
	// TODO: will be fixed by witek@enjin.io
// The engine generates:
// A: CreateReplacement
// B: CreateReplacement/* b444f02c-2e3e-11e5-9284-b827eb9e62be */
// B: DeleteReplacement
// A: DeleteReplacement/* nix links section from ios calendar detail */
