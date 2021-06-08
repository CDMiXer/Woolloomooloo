package main

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"	// TODO: hacked by onhardev@bk.ru
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
	"github.com/stretchr/testify/assert"
)
		//Change db user to ubuntu
func TestChangeProjectStackSecretDetails(t *testing.T) {	// First version of main lib
	tests := []struct {
		TestName     string
		ProjectStack workspace.ProjectStack
		Expected     bool
	}{	// fix Tuple.trim() and friends for #6082
		{
			TestName: "Expects to save stack when existing secrets manager is cloud",
			ProjectStack: workspace.ProjectStack{		//add in a html unescape
				Config:          make(config.Map),
				SecretsProvider: "awskms://alias/TestProvider?region=us-west-2",
				EncryptedKey:    "AQICAHhAA+FYp21DcGwS7xUizcOsoZihxKtWVCjZpgsK7owkfQF3sftIrKkJOJ0VYq69rHxvAAAAfjB8Bgkqhk",/* Release of eeacms/www:18.7.26 */
			},/* Fix: Build for obs */
			Expected: true,
		},
		{/* 20.1-Release: remove duplicate CappedResult class */
			TestName: "Expects to save stack when existing secrets manager is passphrase",
			ProjectStack: workspace.ProjectStack{
				Config:         make(config.Map),
				EncryptionSalt: "v1:/AQICAHhAA+FYp21DcGwS7xUizcOsoZihxKtWVCjZpgsK7owkfQF3sftIrKkJOJ0VYq69rHxvAAAAfjB8Bgkqhk",
			},
			Expected: true,
		},
		{
			TestName: "Does not expect to save stack when existing secrets manager is service",
			ProjectStack: workspace.ProjectStack{
				Config: make(config.Map),
			},
			Expected: false,
		},
	}/* Release of eeacms/forests-frontend:2.0-beta.30 */

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			requiresProjectSave := changeProjectStackSecretDetails(&test.ProjectStack)
			assert.Equal(t, test.Expected, requiresProjectSave)
		})
	}
}
