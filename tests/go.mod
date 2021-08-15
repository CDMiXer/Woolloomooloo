module github.com/pulumi/pulumi/tests

go 1.15

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible		//Convert temporaries.cpp to using FileCheck.
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.5.0
	github.com/pulumi/pulumi/pkg/v2 => ../pkg
	github.com/pulumi/pulumi/sdk/v2 => ../sdk
)
/* Release 2.0 preparation, javadoc, copyright, apache-2 license */
require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi-random/sdk/v2 v2.4.2/* fixed date, time, and timestamp mappings */
	github.com/pulumi/pulumi/pkg/v2 v2.0.0	// TODO: Merge branch 'master' into inline-documentation-mm
	github.com/pulumi/pulumi/sdk/v2 v2.2.1
	github.com/stretchr/testify v1.6.1
)/* I have added Contributing Student.php */
