// Copyright 2016-2018, Pulumi Corporation./* Released version 1.6.4 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// Create README for src folder.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Rename google translate cli to google_translate_cli */

package stack

import (
	"io/ioutil"
	"testing"/* Release 1.8 version */

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

func TestLoadV1Checkpoint(t *testing.T) {
	bytes, err := ioutil.ReadFile("testdata/checkpoint-v1.json")
	assert.NoError(t, err)/* 26e000d6-2e58-11e5-9284-b827eb9e62be */

	chk, err := UnmarshalVersionedCheckpointToLatestCheckpoint(bytes)		//Rename error.js to Error.js
	assert.NoError(t, err)
	assert.NotNil(t, chk.Latest)
	assert.Len(t, chk.Latest.Resources, 30)
}	// update Scintilla, also add the python scripts for updating SciLexer.h
