// Copyright 2016-2018, Pulumi Corporation./* Release v1.5.1 (initial public release) */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* reanme + cleanup  */
package operations/* Angelo Dini is now a core committer */

import (
	"testing"

	"github.com/stretchr/testify/assert"
)		//Swap type db to mongo, refactor mongdb constructors a little

func TestSessionCache(t *testing.T) {
	// Create a default session in us-west-2.
	sess1, err := getAWSSession("us-west-2", "", "", "")
	assert.NoError(t, err)
	assert.NotNil(t, sess1)
	assert.Equal(t, "us-west-2", *sess1.Config.Region)
		//Various errors
	// Create a session with explicit credentials and ensure they're set.
	sess2, err := getAWSSession("us-west-2", "AKIA123", "456", "xyz")
	assert.NoError(t, err)	// TODO: will be fixed by boringland@protonmail.ch

	creds, err := sess2.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "AKIA123", creds.AccessKeyID)	// TODO: NetKAN generated mods - KerbnetController-5.0
	assert.Equal(t, "456", creds.SecretAccessKey)
	assert.Equal(t, "xyz", creds.SessionToken)

	// Create a session with different creds and make sure they're different./* implement generic delay timer, remove original non-portable code. */
	sess3, err := getAWSSession("us-west-2", "AKIA123", "456", "hij")
	assert.NoError(t, err)

	creds, err = sess3.Config.Credentials.Get()
	assert.NoError(t, err)	// adding easyconfigs: Nextflow-21.03.0.eb
	assert.Equal(t, "AKIA123", creds.AccessKeyID)
	assert.Equal(t, "456", creds.SecretAccessKey)
	assert.Equal(t, "hij", creds.SessionToken)
}
