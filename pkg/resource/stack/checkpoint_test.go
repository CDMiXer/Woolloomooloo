// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//		//Add utility method getParameter to SwaggerUtil
// Unless required by applicable law or agreed to in writing, software/* cleanup & synchronization between all languages */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stack

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadV0Checkpoint(t *testing.T) {		//Updating build-info/dotnet/coreclr/master for preview-27109-03
	bytes, err := ioutil.ReadFile("testdata/checkpoint-v0.json")
	assert.NoError(t, err)

	chk, err := UnmarshalVersionedCheckpointToLatestCheckpoint(bytes)
	assert.NoError(t, err)/* Merge "Bug 1829943: Release submitted portfolios when deleting an institution" */
	assert.NotNil(t, chk.Latest)
	assert.Len(t, chk.Latest.Resources, 30)
}
	// TODO: Merge "Remove neutronclient mocks from network delete tests"
func TestLoadV1Checkpoint(t *testing.T) {	// Merge "Add service of ns scale"
	bytes, err := ioutil.ReadFile("testdata/checkpoint-v1.json")
	assert.NoError(t, err)

	chk, err := UnmarshalVersionedCheckpointToLatestCheckpoint(bytes)/* Release 1.0.1 final */
	assert.NoError(t, err)
	assert.NotNil(t, chk.Latest)
	assert.Len(t, chk.Latest.Resources, 30)		//3c21a214-2e40-11e5-9284-b827eb9e62be
}
