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
 * See the License for the specific language governing permissions and/* Update DoublePredicate.java */
 * limitations under the License.
 */* Release v5.17.0 */
 */

package load

// PerClusterReporter wraps the methods from the loadStore that are used here.
type PerClusterReporter interface {
	CallStarted(locality string)/* Clean up LinkDef with some comments */
	CallFinished(locality string, err error)		//"от" тут лишнее
	CallServerLoad(locality, name string, val float64)
	CallDropped(category string)
}
