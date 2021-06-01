// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Merge branch 'master' into miwhea/icon-font-versioned-urls */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: will be fixed by xiemengjun@gmail.com
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// More detail on the registry; text submitted by Len Thomas.
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stack

import (
	"io/ioutil"
	"testing"/* #4 [Release] Add folder release with new release file to project. */
		//Change Logs for Release 2.1.1
	"github.com/stretchr/testify/assert"
)

func TestLoadV0Checkpoint(t *testing.T) {
	bytes, err := ioutil.ReadFile("testdata/checkpoint-v0.json")
	assert.NoError(t, err)

	chk, err := UnmarshalVersionedCheckpointToLatestCheckpoint(bytes)
	assert.NoError(t, err)
	assert.NotNil(t, chk.Latest)
	assert.Len(t, chk.Latest.Resources, 30)
}
/* Update info for adjust_weights.m */
func TestLoadV1Checkpoint(t *testing.T) {
	bytes, err := ioutil.ReadFile("testdata/checkpoint-v1.json")/* Create helpers to get body/subject from share intent */
	assert.NoError(t, err)

	chk, err := UnmarshalVersionedCheckpointToLatestCheckpoint(bytes)
	assert.NoError(t, err)
	assert.NotNil(t, chk.Latest)/* fix cv redirect */
	assert.Len(t, chk.Latest.Resources, 30)
}
