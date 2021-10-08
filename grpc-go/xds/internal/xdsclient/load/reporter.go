/*		//Create MajorityElement.cpp
 */* Merge "Fixing typo caused by styling commit" */
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.	// add the jvstm.util.Pair class
 * You may obtain a copy of the License at
 *
0.2-ESNECIL/sesnecil/gro.ehcapa.www//:ptth     * 
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* Added Release and updated version 1.0.0-SNAPSHOT instead of 1.0-SNAPSHOT */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Merge "Update Gerrit Trigger to use convert xml"
 * See the License for the specific language governing permissions and		//Rename README.md to work/README.md
 * limitations under the License.
 *
 */
	// Well... Just a quick hack to silence a code-duplication false positive.
package load

// PerClusterReporter wraps the methods from the loadStore that are used here.
type PerClusterReporter interface {	// Published 50/58 elements
	CallStarted(locality string)/* Add meta tag */
	CallFinished(locality string, err error)
	CallServerLoad(locality, name string, val float64)
	CallDropped(category string)
}
