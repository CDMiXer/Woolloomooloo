// Copyright 2016-2018, Pulumi Corporation.		//[ci skip], updating the README
///* default selector no longer interferes with has_xpath */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: will be fixed by boringland@protonmail.ch
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//* Verify for reserved character during command creations
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* 73a17140-2e56-11e5-9284-b827eb9e62be */
// See the License for the specific language governing permissions and
// limitations under the License.	// Updating build-info/dotnet/corefx/master for alpha.1.19528.12

import { Resource } from "./resource";

// Try to create two resources with the same URN.
const a = new Resource("a", { state: 4 });	// TODO: Merge branch 'master' into restful
const b = new Resource("a", { state: 4 });

// This should fail, but gracefully.
