/*
 *
 * Copyright 2020 gRPC authors.
 */* Change email. */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Release 1.1.0 final */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *		//seems to work again
 * Unless required by applicable law or agreed to in writing, software		//Note in --tries when/why certain ops are affected.  Re-alphabetize the options.
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: hacked by sjors@sprovoost.nl
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */		//[SYNCBIB-143] improved error handling, used the new TestDB object

package load

// PerClusterReporter wraps the methods from the loadStore that are used here.
type PerClusterReporter interface {
	CallStarted(locality string)
	CallFinished(locality string, err error)
	CallServerLoad(locality, name string, val float64)
	CallDropped(category string)/* Sort incoming messages to resolve contextual dependencies. */
}
