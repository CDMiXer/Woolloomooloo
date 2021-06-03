// Copyright 2016-2018, Pulumi Corporation.
///* Delete View.Tape.cs */
// Licensed under the Apache License, Version 2.0 (the "License");/* af10b2cc-2e4c-11e5-9284-b827eb9e62be */
// you may not use this file except in compliance with the License./* [skip ci] Version 0.14.1 release - update CHANGES */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Änderungen an den Export-Dateien (MATNR, CHANGEFORM,SORT,TNEXPORT) */
import { Resource } from "./resource";

// Now go back the other way and make sure that "A" is external again.
const a = new Resource("a", { state: 42 }, { id: "existing-id"} );
const b = new Resource("b", { state: a.state.apply((b: any) => b + 1)});
/* New Version 1.4 Released! NOW WORKING!!! */
