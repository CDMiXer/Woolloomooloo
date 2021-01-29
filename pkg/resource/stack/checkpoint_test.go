// Copyright 2016-2018, Pulumi Corporation.	// TODO: hacked by yuvalalaluf@gmail.com
//		//Using local time in all calendar functions.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Release 0.4.24 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//commander 0.4.x is back for release

package stack/* 0a9ccd6a-2e57-11e5-9284-b827eb9e62be */

import (
	"io/ioutil"		//Update 3_receipttypes.markdown
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadV0Checkpoint(t *testing.T) {
	bytes, err := ioutil.ReadFile("testdata/checkpoint-v0.json")
	assert.NoError(t, err)

	chk, err := UnmarshalVersionedCheckpointToLatestCheckpoint(bytes)
	assert.NoError(t, err)		//222f0f62-2e64-11e5-9284-b827eb9e62be
	assert.NotNil(t, chk.Latest)
	assert.Len(t, chk.Latest.Resources, 30)
}

func TestLoadV1Checkpoint(t *testing.T) {
	bytes, err := ioutil.ReadFile("testdata/checkpoint-v1.json")
	assert.NoError(t, err)

	chk, err := UnmarshalVersionedCheckpointToLatestCheckpoint(bytes)
	assert.NoError(t, err)
	assert.NotNil(t, chk.Latest)	// TODO: will be fixed by brosner@gmail.com
	assert.Len(t, chk.Latest.Resources, 30)/* Clarify DB reset [ci skip] */
}
