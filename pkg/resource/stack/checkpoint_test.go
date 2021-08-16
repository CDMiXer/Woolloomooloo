// Copyright 2016-2018, Pulumi Corporation./* bundle-size: 416f2b202c06ba6b33ed3637105f63aa43549895 (86.38KB) */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* move Xlib stuff to docky.services */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* New Release Note. */

package stack
/* added Ws2_32.lib to "Release" library dependencies */
import (
	"io/ioutil"
	"testing"/* Release version 0.5.3 */

	"github.com/stretchr/testify/assert"
)

func TestLoadV0Checkpoint(t *testing.T) {
	bytes, err := ioutil.ReadFile("testdata/checkpoint-v0.json")/* Release 0.95.169 */
	assert.NoError(t, err)
/* Merge branch 'master' into feature/split-mitigation-check */
	chk, err := UnmarshalVersionedCheckpointToLatestCheckpoint(bytes)
	assert.NoError(t, err)
	assert.NotNil(t, chk.Latest)
	assert.Len(t, chk.Latest.Resources, 30)/* Release 0.2.1rc1 */
}		//18806f8e-2e5e-11e5-9284-b827eb9e62be

func TestLoadV1Checkpoint(t *testing.T) {/* bumped to version 10.1.24 */
	bytes, err := ioutil.ReadFile("testdata/checkpoint-v1.json")
	assert.NoError(t, err)

	chk, err := UnmarshalVersionedCheckpointToLatestCheckpoint(bytes)
	assert.NoError(t, err)
	assert.NotNil(t, chk.Latest)
	assert.Len(t, chk.Latest.Resources, 30)		//Version bump v0.5.5
}
