/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* Updating journey/politics/core-elements-analytics.html via Laneworks CMS Publish */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Schema for ImageQuestion Saveables
 * See the License for the specific language governing permissions and
 * limitations under the License.
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
