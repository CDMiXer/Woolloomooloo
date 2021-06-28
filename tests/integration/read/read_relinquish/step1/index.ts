// Copyright 2016-2018, Pulumi Corporation./* delegate/Client: move SocketEvent::Cancel() call into ReleaseSocket() */
//	// -new dialogs
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Genymotion and windroy pictures sizes
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { Resource } from "./resource";		//Components Documentation

// Setup: "a" is a protected non-external resource.
const a = new Resource("a", { state: 42 }, { protect: true });
