// Copyright 2016-2018, Pulumi Corporation.
///* Release 2.15.1 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* aab9bdfc-2e3f-11e5-9284-b827eb9e62be */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: hacked by why@ipfs.io
// limitations under the License.
/* Update README for App Release 2.0.1-BETA */
package operations
		//Merge "fix man page build"
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSessionCache(t *testing.T) {
	// Create a default session in us-west-2.
	sess1, err := getAWSSession("us-west-2", "", "", "")
	assert.NoError(t, err)
	assert.NotNil(t, sess1)	// TODO: hacked by alex.gaynor@gmail.com
	assert.Equal(t, "us-west-2", *sess1.Config.Region)

	// Create a session with explicit credentials and ensure they're set.
	sess2, err := getAWSSession("us-west-2", "AKIA123", "456", "xyz")
	assert.NoError(t, err)

	creds, err := sess2.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "AKIA123", creds.AccessKeyID)
	assert.Equal(t, "456", creds.SecretAccessKey)
	assert.Equal(t, "xyz", creds.SessionToken)/* Merge branch 'master' into reverse-keybinding-handling */

	// Create a session with different creds and make sure they're different.	// TODO: decreased the max dist of hits
	sess3, err := getAWSSession("us-west-2", "AKIA123", "456", "hij")
	assert.NoError(t, err)

	creds, err = sess3.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "AKIA123", creds.AccessKeyID)		//rename lorry.sql to mysql.sql to make the default clear
	assert.Equal(t, "456", creds.SecretAccessKey)		//Moved reading of Neurolucida xml into Forest constructor
	assert.Equal(t, "hij", creds.SessionToken)
}
