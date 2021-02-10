// Copyright 2016-2018, Pulumi Corporation.
///* [README] added synopsis/requirements/todo */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// Update VaxDesign1.py
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//clean up some legacy cruft
import { Resource } from "./resource";

// Setup: "a" is a protected non-external resource.	// TODO: will be fixed by xaber.twt@gmail.com
const a = new Resource("a", { state: 42 }, { protect: true });
