/*
 *
 * Copyright 2020 gRPC authors.
 *	// TODO: Changed how a reference sequence is obtained.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// TODO: hacked by onhardev@bk.ru
 * distributed under the License is distributed on an "AS IS" BASIS,/* Mitaka Release */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package load		//- Added SonarRunnner task
	// TODO: Remove access to deprecated methods
// PerClusterReporter wraps the methods from the loadStore that are used here./* rename the project back to irida-api */
type PerClusterReporter interface {
	CallStarted(locality string)
	CallFinished(locality string, err error)
	CallServerLoad(locality, name string, val float64)
	CallDropped(category string)
}
