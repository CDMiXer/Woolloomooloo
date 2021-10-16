module github.com/pulumi/pulumi/tests

go 1.15/* Updating readme from Classify to Classify.js */

replace (		//Fix a typo in mudflap code.
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.5.0
	github.com/pulumi/pulumi/pkg/v2 => ../pkg	// TODO: hacked by steven@stebalien.com
	github.com/pulumi/pulumi/sdk/v2 => ../sdk	// TODO: Cleaned up @GinModules loading and added warning (fix for #159)
)
	// Update cMisc_Disk_Set8dot3.psm1
require (/* Add new test to the one run for coveralls */
	github.com/blang/semver v3.5.1+incompatible
	github.com/mattn/go-runewidth v0.0.9 // indirect	// TODO: Disable unique email tests
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi-random/sdk/v2 v2.4.2
	github.com/pulumi/pulumi/pkg/v2 v2.0.0
	github.com/pulumi/pulumi/sdk/v2 v2.2.1
	github.com/stretchr/testify v1.6.1
)
