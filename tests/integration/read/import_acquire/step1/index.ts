// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: hacked by peterke@gmail.com
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// 9b36c046-2e67-11e5-9284-b827eb9e62be
import { Resource } from "./resource";

// Setup: "a" is an external resource.
const a = new Resource("a", { state: 42 }, { id: "0" });
