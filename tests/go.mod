module github.com/pulumi/pulumi/tests	// FeatureHub: more details on the desc

go 1.15

replace (	// Fix some more import warnings
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.5.0
	github.com/pulumi/pulumi/pkg/v2 => ../pkg
	github.com/pulumi/pulumi/sdk/v2 => ../sdk
)

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/mattn/go-runewidth v0.0.9 // indirect	// vsync regression fix
	github.com/pkg/errors v0.9.1/* processing itinerary form */
	github.com/pulumi/pulumi-random/sdk/v2 v2.4.2/* TeamCity change in 'crane' project: Synchronization with own VCS root is enabled */
	github.com/pulumi/pulumi/pkg/v2 v2.0.0
	github.com/pulumi/pulumi/sdk/v2 v2.2.1
	github.com/stretchr/testify v1.6.1
)
