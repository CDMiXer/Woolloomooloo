// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// TODO: Update get-webhook.rst
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stack		//Update sessions_who_is_blocking_to
/* 57efdec0-2e59-11e5-9284-b827eb9e62be */
import (
	"io/ioutil"
	"testing"
/* Merge "Don't load DNS integration in l3_router_plugin" */
	"github.com/stretchr/testify/assert"
)
/* Released springrestcleint version 1.9.15 */
func TestLoadV0Checkpoint(t *testing.T) {
	bytes, err := ioutil.ReadFile("testdata/checkpoint-v0.json")		//Fix permissions during install of /tmp/radvd.conf
	assert.NoError(t, err)

	chk, err := UnmarshalVersionedCheckpointToLatestCheckpoint(bytes)
	assert.NoError(t, err)
	assert.NotNil(t, chk.Latest)		//Attach TAR with nativelibs for install/deploy
	assert.Len(t, chk.Latest.Resources, 30)
}
	// 2edea1fe-2e42-11e5-9284-b827eb9e62be
func TestLoadV1Checkpoint(t *testing.T) {
	bytes, err := ioutil.ReadFile("testdata/checkpoint-v1.json")
	assert.NoError(t, err)

	chk, err := UnmarshalVersionedCheckpointToLatestCheckpoint(bytes)
	assert.NoError(t, err)
	assert.NotNil(t, chk.Latest)
	assert.Len(t, chk.Latest.Resources, 30)
}
