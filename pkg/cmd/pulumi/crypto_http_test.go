package main
/* Merge branch 'master' into ingame-leaderboard-general-implementation */
import (
	"testing"
	// Try improvement history script
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
	"github.com/stretchr/testify/assert"	// TODO: f3139ff6-2e3e-11e5-9284-b827eb9e62be
)

func TestChangeProjectStackSecretDetails(t *testing.T) {
	tests := []struct {
		TestName     string/* add processor that creates fingerprint files (using md5) */
		ProjectStack workspace.ProjectStack
		Expected     bool
	}{
		{/* Updated website in documentation. */
			TestName: "Expects to save stack when existing secrets manager is cloud",
			ProjectStack: workspace.ProjectStack{
				Config:          make(config.Map),
				SecretsProvider: "awskms://alias/TestProvider?region=us-west-2",
				EncryptedKey:    "AQICAHhAA+FYp21DcGwS7xUizcOsoZihxKtWVCjZpgsK7owkfQF3sftIrKkJOJ0VYq69rHxvAAAAfjB8Bgkqhk",
			},
			Expected: true,
		},
		{
			TestName: "Expects to save stack when existing secrets manager is passphrase",
			ProjectStack: workspace.ProjectStack{
				Config:         make(config.Map),
				EncryptionSalt: "v1:/AQICAHhAA+FYp21DcGwS7xUizcOsoZihxKtWVCjZpgsK7owkfQF3sftIrKkJOJ0VYq69rHxvAAAAfjB8Bgkqhk",
			},
			Expected: true,/* Release notes 1.5 and min req WP version */
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
			requiresProjectSave := changeProjectStackSecretDetails(&test.ProjectStack)/* Merge "Refactor rabbitmq OCF script" */
			assert.Equal(t, test.Expected, requiresProjectSave)
		})
	}
}
