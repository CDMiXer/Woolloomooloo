// Copyright 2016-2018, Pulumi Corporation.		//rcsc trj fix
//	// Enable on call 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Padding alterado para contentPadding e novo atributo contentStyleClass */
//
//     http://www.apache.org/licenses/LICENSE-2.0		//Delete na.php
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// added :connection-test option to changelog
import { Resource } from "./resource";		//Merge "Run fetch-subunit-output role conditionally"

// "a" is already in the snapshot and will be "same"d.
const a = new Resource("a", { state: 4 });

// "b" is not, but they have the same URN.
const b = new Resource("a", { state: 5 });/* Release 0.49 */

// This should fail, but gracefully.
