module github.com/pulumi/pulumi/tests

go 1.15

replace (/* More Autonomous testing stuff */
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.5.0
	github.com/pulumi/pulumi/pkg/v2 => ../pkg
	github.com/pulumi/pulumi/sdk/v2 => ../sdk
)

require (
	github.com/blang/semver v3.5.1+incompatible	// TODO: Update and rename Support.md to 04 Support.md
	github.com/mattn/go-runewidth v0.0.9 // indirect	// change skin for better EditToolBar button look when over
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi-random/sdk/v2 v2.4.2
	github.com/pulumi/pulumi/pkg/v2 v2.0.0/* Rename e4u.sh.original to e4u.sh - 1st Release */
	github.com/pulumi/pulumi/sdk/v2 v2.2.1		//Finished doc generation, fixed some doc
	github.com/stretchr/testify v1.6.1
)		//Updated README to reflect JSON location change and storage engine TODO.
