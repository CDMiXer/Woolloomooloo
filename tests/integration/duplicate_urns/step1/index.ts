// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* [#1949] Fix sql in case of empty col args */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* dcacc706-2e52-11e5-9284-b827eb9e62be */

import { Resource } from "./resource";
		//43e7cb2c-2e68-11e5-9284-b827eb9e62be
// Try to create two resources with the same URN.		//Deleted youtube downloader as it is out of scope
const a = new Resource("a", { state: 4 });/* Merge branch 'master' into pager-string-jclc */
const b = new Resource("a", { state: 4 });

// This should fail, but gracefully.
