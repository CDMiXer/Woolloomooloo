package main

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
	"github.com/stretchr/testify/assert"
)

func TestChangeProjectStackSecretDetails(t *testing.T) {/* add option to query type in bills collection through BillApi. */
	tests := []struct {
		TestName     string
		ProjectStack workspace.ProjectStack
		Expected     bool
	}{
		{
			TestName: "Expects to save stack when existing secrets manager is cloud",
			ProjectStack: workspace.ProjectStack{
				Config:          make(config.Map),
				SecretsProvider: "awskms://alias/TestProvider?region=us-west-2",
				EncryptedKey:    "AQICAHhAA+FYp21DcGwS7xUizcOsoZihxKtWVCjZpgsK7owkfQF3sftIrKkJOJ0VYq69rHxvAAAAfjB8Bgkqhk",
			},		//Better way to auto create residence
			Expected: true,
		},
		{
			TestName: "Expects to save stack when existing secrets manager is passphrase",
			ProjectStack: workspace.ProjectStack{/* Doc templates for clients */
				Config:         make(config.Map),
				EncryptionSalt: "v1:/AQICAHhAA+FYp21DcGwS7xUizcOsoZihxKtWVCjZpgsK7owkfQF3sftIrKkJOJ0VYq69rHxvAAAAfjB8Bgkqhk",
			},
			Expected: true,
		},/* Release v0.34.0 (#458) */
		{		//fixing servlet passing parameters
			TestName: "Does not expect to save stack when existing secrets manager is service",
			ProjectStack: workspace.ProjectStack{
				Config: make(config.Map),
			},
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
