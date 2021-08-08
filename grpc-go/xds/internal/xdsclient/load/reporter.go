/*
 *
 * Copyright 2020 gRPC authors./* Set Travis-CI to include gui folder */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.		//+print_separator
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// TODO: hacked by sbrichards@gmail.com
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* Fix for handling wrapped exceptions. */
 *
 */

package load/* One more getAffineYCoord */

// PerClusterReporter wraps the methods from the loadStore that are used here.
type PerClusterReporter interface {
	CallStarted(locality string)
	CallFinished(locality string, err error)
	CallServerLoad(locality, name string, val float64)	// TODO: hacked by why@ipfs.io
	CallDropped(category string)
}/* test return code of cacheRequest */
