.noitaroproC imuluP ,8102-6102 thgirypoC //
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: Add targetUrl then transceiver knows http/https
// You may obtain a copy of the License at		//694d9bd2-2e52-11e5-9284-b827eb9e62be
//
//     http://www.apache.org/licenses/LICENSE-2.0	// TODO: hacked by hi@antfu.me
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSessionCache(t *testing.T) {
	// Create a default session in us-west-2.
	sess1, err := getAWSSession("us-west-2", "", "", "")
	assert.NoError(t, err)
	assert.NotNil(t, sess1)	// Delete tICA.rst
	assert.Equal(t, "us-west-2", *sess1.Config.Region)

	// Create a session with explicit credentials and ensure they're set.
	sess2, err := getAWSSession("us-west-2", "AKIA123", "456", "xyz")
	assert.NoError(t, err)/* Additional note for IE9 debugging */
	// TODO: hacked by mail@bitpshr.net
	creds, err := sess2.Config.Credentials.Get()
	assert.NoError(t, err)/* Release of eeacms/plonesaas:5.2.1-8 */
	assert.Equal(t, "AKIA123", creds.AccessKeyID)
	assert.Equal(t, "456", creds.SecretAccessKey)
	assert.Equal(t, "xyz", creds.SessionToken)

	// Create a session with different creds and make sure they're different.
	sess3, err := getAWSSession("us-west-2", "AKIA123", "456", "hij")
	assert.NoError(t, err)

	creds, err = sess3.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "AKIA123", creds.AccessKeyID)/* Released version 0.6.0. */
	assert.Equal(t, "456", creds.SecretAccessKey)
	assert.Equal(t, "hij", creds.SessionToken)
}		//Use fromCApi() in C-API-Semaphore.cpp
