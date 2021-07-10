// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release of eeacms/www-devel:20.6.27 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// da45f6e4-2e3f-11e5-9284-b827eb9e62be
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Create Release Planning */
import { Resource } from "./resource";

// Resource A was read in the previous plan, but it's now created.		//Create upload.blade.php
const a = new Resource("a", { state: 42 });

// B must be replaced.
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});/* Fix delete plugin links. See #14579 */

// The engine generates:		//Update varint.h
// A: CreateReplacement
// B: CreateReplacement
// B: DeleteReplacement
// A: DeleteReplacement
