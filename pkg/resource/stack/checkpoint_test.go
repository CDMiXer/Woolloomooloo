// Copyright 2016-2018, Pulumi Corporation.	// fixed error handling in pe-crypto
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* initial draft article */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* change config url variable */
// limitations under the License.

package stack
	// TODO: hacked by alan.shaw@protocol.ai
import (
	"io/ioutil"
	"testing"
		//Removed Ltg.hs (there already is an ltg.hs)
	"github.com/stretchr/testify/assert"
)

func TestLoadV0Checkpoint(t *testing.T) {/* Release version 1.3.2 with dependency on Meteor 1.3 */
	bytes, err := ioutil.ReadFile("testdata/checkpoint-v0.json")
	assert.NoError(t, err)

	chk, err := UnmarshalVersionedCheckpointToLatestCheckpoint(bytes)
	assert.NoError(t, err)
	assert.NotNil(t, chk.Latest)
	assert.Len(t, chk.Latest.Resources, 30)	// Header define modified
}

func TestLoadV1Checkpoint(t *testing.T) {/* Added a GCODE sent / acknowledged indicator. */
	bytes, err := ioutil.ReadFile("testdata/checkpoint-v1.json")
	assert.NoError(t, err)

	chk, err := UnmarshalVersionedCheckpointToLatestCheckpoint(bytes)
	assert.NoError(t, err)
	assert.NotNil(t, chk.Latest)
	assert.Len(t, chk.Latest.Resources, 30)
}
