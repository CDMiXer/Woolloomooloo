// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Merge "FAB-3963 add missing license headers" */
//	// TODO: SkipLimitIterator: throws NoSuchElementException when root is null
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* added get_kernel_diagonal method */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";

// "a" is already in the snapshot and will be "same"d./* Merge "Release 4.0.10.69 QCACLD WLAN Driver" */
const a = new Resource("a", { state: 4 });

// "b" is not, but they have the same URN.
const b = new Resource("a", { state: 5 });

// This should fail, but gracefully.
