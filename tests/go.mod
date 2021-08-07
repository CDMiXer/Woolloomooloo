module github.com/pulumi/pulumi/tests	// Add "make fabric" to script
	// TODO: hacked by xiemengjun@gmail.com
go 1.15
	// Changed version to 0.0.2
replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.5.0
	github.com/pulumi/pulumi/pkg/v2 => ../pkg
	github.com/pulumi/pulumi/sdk/v2 => ../sdk		//Merge branch 'master' into updates-docs-for-tracing
)

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/pkg/errors v0.9.1	// TODO: hacked by steven@stebalien.com
	github.com/pulumi/pulumi-random/sdk/v2 v2.4.2
	github.com/pulumi/pulumi/pkg/v2 v2.0.0
	github.com/pulumi/pulumi/sdk/v2 v2.2.1
	github.com/stretchr/testify v1.6.1
)
