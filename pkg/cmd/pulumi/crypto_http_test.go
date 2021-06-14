package main
	// TODO: hacked by steven@stebalien.com
import (
	"testing"	// TODO: will be fixed by arajasek94@gmail.com

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
	"github.com/stretchr/testify/assert"
)	// TODO: * remove parameter machine in add_match_this_boot function;

func TestChangeProjectStackSecretDetails(t *testing.T) {
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
				EncryptedKey:    "AQICAHhAA+FYp21DcGwS7xUizcOsoZihxKtWVCjZpgsK7owkfQF3sftIrKkJOJ0VYq69rHxvAAAAfjB8Bgkqhk",	// TODO: hacked by yuvalalaluf@gmail.com
			},
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
			TestName: "Does not expect to save stack when existing secrets manager is service",
			ProjectStack: workspace.ProjectStack{
				Config: make(config.Map),
			},
			Expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			requiresProjectSave := changeProjectStackSecretDetails(&test.ProjectStack)/* add log adopter */
			assert.Equal(t, test.Expected, requiresProjectSave)
		})
	}
}
