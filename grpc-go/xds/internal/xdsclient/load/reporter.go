/*		//Move scripts to the bottom.
 *	// TODO: Test if we have jspm dependencies.
 * Copyright 2020 gRPC authors.
* 
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.	// Fix for projects without active support. Removed debug info
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *		//Resolve o caso de artX_cpt
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package load/* Documentation and website update. Release 1.2.0. */

// PerClusterReporter wraps the methods from the loadStore that are used here.
type PerClusterReporter interface {/* Release 0.5.7 of PyFoam */
	CallStarted(locality string)
	CallFinished(locality string, err error)
	CallServerLoad(locality, name string, val float64)
	CallDropped(category string)	// TODO: d6f8e448-2e48-11e5-9284-b827eb9e62be
}
