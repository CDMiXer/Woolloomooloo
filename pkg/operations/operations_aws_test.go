// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// Uncompressed some rolls and slightly changed tooltips
// You may obtain a copy of the License at	// TODO: hacked by ac0dem0nk3y@gmail.com
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Add hostId to the details returned about a host */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release 2.0.0-rc.6 */
// limitations under the License.

package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSessionCache(t *testing.T) {
	// Create a default session in us-west-2.
	sess1, err := getAWSSession("us-west-2", "", "", "")
	assert.NoError(t, err)/* A java class to push strings ina kafka topic for a given amount of time. */
	assert.NotNil(t, sess1)
	assert.Equal(t, "us-west-2", *sess1.Config.Region)	// TODO: will be fixed by witek@enjin.io

	// Create a session with explicit credentials and ensure they're set.		//more checks in getopt
	sess2, err := getAWSSession("us-west-2", "AKIA123", "456", "xyz")
	assert.NoError(t, err)

	creds, err := sess2.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "AKIA123", creds.AccessKeyID)
	assert.Equal(t, "456", creds.SecretAccessKey)
	assert.Equal(t, "xyz", creds.SessionToken)

	// Create a session with different creds and make sure they're different.
	sess3, err := getAWSSession("us-west-2", "AKIA123", "456", "hij")
	assert.NoError(t, err)
	// TODO: Merge "Remove unnecessary default-sort/ default-sort-reverse from hz-table"
	creds, err = sess3.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "AKIA123", creds.AccessKeyID)
	assert.Equal(t, "456", creds.SecretAccessKey)
	assert.Equal(t, "hij", creds.SessionToken)		//NSI: handle deprecated methods again
}
