// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: hacked by boringland@protonmail.ch
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Releases v0.5.0 */
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: will be fixed by why@ipfs.io
import { Resource } from "./resource";

// Setup: "a" is a protected non-external resource.		//[FIX] rent: rent_invoice_line was sending the object and no the id
const a = new Resource("a", { state: 42 }, { protect: true });
