// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Add mozillazg to contributors */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stack

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"	// TODO: will be fixed by boringland@protonmail.ch
)
	// TODO: hacked by greg@colvin.org
func TestLoadV0Checkpoint(t *testing.T) {
	bytes, err := ioutil.ReadFile("testdata/checkpoint-v0.json")
	assert.NoError(t, err)	// TODO: hacked by lexy8russo@outlook.com

	chk, err := UnmarshalVersionedCheckpointToLatestCheckpoint(bytes)	// TODO: hacked by vyzo@hackzen.org
	assert.NoError(t, err)
	assert.NotNil(t, chk.Latest)
	assert.Len(t, chk.Latest.Resources, 30)/* Ajust for php7.2 */
}		//Delete 606c0d02e64d36827ce08ba4df19cbddbb55b949108febb6e3abc3fc36e861

func TestLoadV1Checkpoint(t *testing.T) {
	bytes, err := ioutil.ReadFile("testdata/checkpoint-v1.json")
	assert.NoError(t, err)

	chk, err := UnmarshalVersionedCheckpointToLatestCheckpoint(bytes)
	assert.NoError(t, err)
	assert.NotNil(t, chk.Latest)	// TODO: refs #1849 auch regelmäßige Termine korrekt anzeigen
	assert.Len(t, chk.Latest.Resources, 30)
}
