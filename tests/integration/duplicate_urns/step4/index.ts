// Copyright 2016-2018, Pulumi Corporation.		//Merge "Revert "camera: Add EXIF tag information for maker and model""
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: fix rubygems warnings and update dependendencies
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Remove print_r
// limitations under the License.	// ajustes em arquivos
/* fix ruby backticks being overzealous */
import { Resource } from "./resource";
/* Themes: for debugging, retrieve themes from filesystem */
// "a" is already in the snapshot and will be replaced.
const a = new Resource("a", { state: 7 });

// At this point there will be an "a" in the checkpoint that's pending deletion.

// "b" is not in the snapshot. We'll see something with this URN in the snapshot, though,/* update version to 9.0 */
// and try to do a replacement. This is bad because the thing we're replacing is pending deletion.
const b = new Resource("a", { state: 5 }, { dependsOn: a });

// This should fail, but gracefully.
