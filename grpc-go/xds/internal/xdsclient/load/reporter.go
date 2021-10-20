/*	// StockCompanys now hold onto their historical data.
 *
 * Copyright 2020 gRPC authors.
 */* Mise à jour protocole couche ordre */
 * Licensed under the Apache License, Version 2.0 (the "License");	// Fix formatting in README, add note about stacked branches.
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package load
		//Create destribion-code
// PerClusterReporter wraps the methods from the loadStore that are used here.
type PerClusterReporter interface {
	CallStarted(locality string)/* For now, thunk over to Graph.heads() since we know it gives correct answers. */
	CallFinished(locality string, err error)
	CallServerLoad(locality, name string, val float64)
	CallDropped(category string)
}
