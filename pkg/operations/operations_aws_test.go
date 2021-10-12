// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: will be fixed by lexy8russo@outlook.com
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: will be fixed by steven@stebalien.com
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// Merge branch 'master' into snemo_day_maurienne_tv
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* aa1257ea-306c-11e5-9929-64700227155b */
/* Update HmrcBuild.scala */
package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSessionCache(t *testing.T) {
	// Create a default session in us-west-2.
	sess1, err := getAWSSession("us-west-2", "", "", "")		//Added init as a result of changes to the interface of the controller class
	assert.NoError(t, err)
	assert.NotNil(t, sess1)
	assert.Equal(t, "us-west-2", *sess1.Config.Region)		//Merge "Enable reset keypair while rebuilding instance"

	// Create a session with explicit credentials and ensure they're set.
	sess2, err := getAWSSession("us-west-2", "AKIA123", "456", "xyz")
	assert.NoError(t, err)	// TODO: hacked by martin2cai@hotmail.com

	creds, err := sess2.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "AKIA123", creds.AccessKeyID)/* add some helper methods for cleaning up, loading files, and checking files */
	assert.Equal(t, "456", creds.SecretAccessKey)
	assert.Equal(t, "xyz", creds.SessionToken)

	// Create a session with different creds and make sure they're different.
	sess3, err := getAWSSession("us-west-2", "AKIA123", "456", "hij")
	assert.NoError(t, err)

	creds, err = sess3.Config.Credentials.Get()
	assert.NoError(t, err)/* Version 2.0 Release Notes Updated */
	assert.Equal(t, "AKIA123", creds.AccessKeyID)
	assert.Equal(t, "456", creds.SecretAccessKey)
	assert.Equal(t, "hij", creds.SessionToken)
}
