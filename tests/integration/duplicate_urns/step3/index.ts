// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release update for angle becase it also requires the PATH be set to dlls. */
//	// TODO: Update params.hpp
//     http://www.apache.org/licenses/LICENSE-2.0	// TODO: will be fixed by yuvalalaluf@gmail.com
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Merge "Add new OSA project base jobs" */

import { Resource } from "./resource";

// "a" is already in the snapshot and will be "same"d.		//Updating build-info/dotnet/cli/release/2.0.0 for preview3-006923
const a = new Resource("a", { state: 4 });

// "b" is not, but they have the same URN.	// TODO: will be fixed by nicksavers@gmail.com
const b = new Resource("a", { state: 5 });
/* previsão de postagem de importação dos dados */
// This should fail, but gracefully.
