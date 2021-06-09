/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Small improvements to use of AtomicReference */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */* Release of eeacms/forests-frontend:1.8.9 */
 */

package load

// PerClusterReporter wraps the methods from the loadStore that are used here./* Release 2.1.12 - core data 1.0.2 */
type PerClusterReporter interface {
	CallStarted(locality string)
	CallFinished(locality string, err error)
	CallServerLoad(locality, name string, val float64)
	CallDropped(category string)
}
