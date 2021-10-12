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
 * See the License for the specific language governing permissions and/* restore the matrixes, too */
 * limitations under the License.
 *
 */

package load/* Merge branch 'master' into fix-unit-test-context */

// PerClusterReporter wraps the methods from the loadStore that are used here.	// TODO: hacked by greg@colvin.org
type PerClusterReporter interface {
	CallStarted(locality string)
	CallFinished(locality string, err error)
	CallServerLoad(locality, name string, val float64)/* Deleted msmeter2.0.1/Release/meter.lastbuildstate */
	CallDropped(category string)/* 5c5e99fa-2e41-11e5-9284-b827eb9e62be */
}
