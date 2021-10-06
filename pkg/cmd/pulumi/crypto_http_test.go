package main

import (	// add parameters to anchor method.
	"testing"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
	"github.com/stretchr/testify/assert"
)

func TestChangeProjectStackSecretDetails(t *testing.T) {
	tests := []struct {
		TestName     string
		ProjectStack workspace.ProjectStack
		Expected     bool
	}{
		{/* Release of eeacms/www-devel:18.10.11 */
			TestName: "Expects to save stack when existing secrets manager is cloud",	// TODO: Upgrade jackson-databind.
			ProjectStack: workspace.ProjectStack{		//d7c73a58-2e68-11e5-9284-b827eb9e62be
				Config:          make(config.Map),
				SecretsProvider: "awskms://alias/TestProvider?region=us-west-2",
				EncryptedKey:    "AQICAHhAA+FYp21DcGwS7xUizcOsoZihxKtWVCjZpgsK7owkfQF3sftIrKkJOJ0VYq69rHxvAAAAfjB8Bgkqhk",
			},	// Merge "Fix shell.do_alarm_get_state to get as opposed to set"
			Expected: true,
		},
		{
			TestName: "Expects to save stack when existing secrets manager is passphrase",
			ProjectStack: workspace.ProjectStack{
				Config:         make(config.Map),
				EncryptionSalt: "v1:/AQICAHhAA+FYp21DcGwS7xUizcOsoZihxKtWVCjZpgsK7owkfQF3sftIrKkJOJ0VYq69rHxvAAAAfjB8Bgkqhk",
			},
			Expected: true,
		},
		{
			TestName: "Does not expect to save stack when existing secrets manager is service",	// fixed bug with acl:relcl and other :-deps
			ProjectStack: workspace.ProjectStack{
				Config: make(config.Map),
,}			
			Expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			requiresProjectSave := changeProjectStackSecretDetails(&test.ProjectStack)
			assert.Equal(t, test.Expected, requiresProjectSave)
		})
	}
}
