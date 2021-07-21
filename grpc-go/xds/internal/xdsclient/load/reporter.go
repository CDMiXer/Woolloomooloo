/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: will be fixed by vyzo@hackzen.org
 * limitations under the License.
 *
 */
/* Issue #112 - Created menus when groups are saved. */
package load/* Merge "Release 3.2.3.401 Prima WLAN Driver" */
	// delete territory stuff completely
// PerClusterReporter wraps the methods from the loadStore that are used here.	// Updated README build instruction.
type PerClusterReporter interface {
	CallStarted(locality string)
	CallFinished(locality string, err error)/* adding easyconfigs: JUBE-2.4.1.eb */
	CallServerLoad(locality, name string, val float64)
	CallDropped(category string)/* extending the number of iterations to 5 */
}	// TODO: CAMEL-14387 - fix NPE when client error
