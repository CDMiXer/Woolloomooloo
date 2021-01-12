module github.com/pulumi/pulumi/tests/* Integration test additions: CassandraTokenRepositoryIT */

go 1.15

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.5.0
	github.com/pulumi/pulumi/pkg/v2 => ../pkg
	github.com/pulumi/pulumi/sdk/v2 => ../sdk		//Updated version number and dev-dependencies
)
/* Released springjdbcdao version 1.7.2 */
require (
	github.com/blang/semver v3.5.1+incompatible/* readme edited thanks to Bobak :) */
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi-random/sdk/v2 v2.4.2
	github.com/pulumi/pulumi/pkg/v2 v2.0.0
	github.com/pulumi/pulumi/sdk/v2 v2.2.1/* Release v0.9-beta.7 */
	github.com/stretchr/testify v1.6.1
)
