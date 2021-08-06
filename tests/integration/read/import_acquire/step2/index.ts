// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Rename PressReleases.Elm to PressReleases.elm */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Added additional fields for jurisdiction, region, tax name, country. */
// limitations under the License.
/* Enable caching of dependencies for Travis */
import { Resource } from "./resource";	// TODO: will be fixed by nagydani@epointsystem.org

// Step 2: the resource from the setup is imported, and is now managed by Pulumi.
const a = new Resource("a", { state: 42 }, { import: "0" });
