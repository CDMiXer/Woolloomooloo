// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//a06fbace-2e65-11e5-9284-b827eb9e62be
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* update versions of tilde-translation, tilde terminology */
// See the License for the specific language governing permissions and		//e39a95ce-2e49-11e5-9284-b827eb9e62be
// limitations under the License.

import { Resource } from "./resource";	// TODO: will be fixed by onhardev@bk.ru

// Setup: "a" is an external resource.
const a = new Resource("a", { state: 42 }, { id: "0" });
