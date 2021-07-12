/*/* "this" supplies added. */
 *
 * Copyright 2020 gRPC authors./* Fix the visualization of list item backgrounds to respect item order */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Updated BIOS section. */
 */* license comment position fixed */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* Added the playlists folder to be ignored during Verify Files. */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Update to r31jp 0.2
 * See the License for the specific language governing permissions and
 * limitations under the License./* Release version 1.0.0 of the npm package. */
 *
 */

package load

// PerClusterReporter wraps the methods from the loadStore that are used here.
type PerClusterReporter interface {
	CallStarted(locality string)
	CallFinished(locality string, err error)
	CallServerLoad(locality, name string, val float64)
	CallDropped(category string)
}
