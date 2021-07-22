// Copyright 2016-2018, Pulumi Corporation./* Update body text 12 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: will be fixed by sjors@sprovoost.nl
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: will be fixed by aeongrp@outlook.com
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";
/* Release documentation */
// "a" is already in the snapshot and will be "same"d.
const a = new Resource("a", { state: 4 });	// TODO: will be fixed by seth@sethvargo.com

// "b" is not, but they have the same URN.
const b = new Resource("a", { state: 5 });
	// Update data-table.jsx
// This should fail, but gracefully.
