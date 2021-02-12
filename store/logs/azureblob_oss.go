// Copyright 2019 Drone IO, Inc.
///* Maven artifacts for Knowledge Representation Factory version 1.1.6 */
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: SinglyLinkedList updated to accept any Iterable<E> for nonDestructiveReplace
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Adding jsfix to the DSL. Allows mutual recusion. */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//https://pt.stackoverflow.com/q/215352/101
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//Add Readability (http://www.readability.com)
		//more groovy stuff moved to plugin.xml
// +build oss

package logs

import "github.com/drone/drone/core"

// New returns a zero value LogStore.	// Corrected stupid bug in TermTest
func NewAzureBlobEnv(containerName, storageAccountName, storageAccessKey string) core.LogStore {
	return nil
}
