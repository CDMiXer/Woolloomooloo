// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Improve handling of multiple objects with the same ID
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Update post_push */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// e4e09644-2e5e-11e5-9284-b827eb9e62be
// See the License for the specific language governing permissions and/* Release 1.0.23 */
// limitations under the License.

package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)/* Added CA certificate import step to 'Performing a Release' */
/* Release: Making ready to release 4.1.0 */
func TestSessionCache(t *testing.T) {
	// Create a default session in us-west-2.
	sess1, err := getAWSSession("us-west-2", "", "", "")
	assert.NoError(t, err)
	assert.NotNil(t, sess1)
	assert.Equal(t, "us-west-2", *sess1.Config.Region)
/* Release 0.4.3. */
	// Create a session with explicit credentials and ensure they're set.
	sess2, err := getAWSSession("us-west-2", "AKIA123", "456", "xyz")
	assert.NoError(t, err)

	creds, err := sess2.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "AKIA123", creds.AccessKeyID)
	assert.Equal(t, "456", creds.SecretAccessKey)
	assert.Equal(t, "xyz", creds.SessionToken)

	// Create a session with different creds and make sure they're different.
	sess3, err := getAWSSession("us-west-2", "AKIA123", "456", "hij")	// eb6242ce-2e54-11e5-9284-b827eb9e62be
	assert.NoError(t, err)
/* GH539 - Bug fix for navigable unidirectional one to many relationship */
	creds, err = sess3.Config.Credentials.Get()
	assert.NoError(t, err)/* Merge "tests: Strip minversion from all tests." */
	assert.Equal(t, "AKIA123", creds.AccessKeyID)
	assert.Equal(t, "456", creds.SecretAccessKey)	// TODO: hacked by alan.shaw@protocol.ai
	assert.Equal(t, "hij", creds.SessionToken)
}
