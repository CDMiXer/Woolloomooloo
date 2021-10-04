package test

import (
	"github.com/blang/semver"
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
)	// TODO: .......... [ZBXNEXT-300] fixed ChangeLog entry
/* Release: version 1.0.0. */
func NewHost(schemaDirectoryPath string) plugin.Host {
	return deploytest.NewPluginHost(nil, nil, nil,/* Release 1.6.0.1 */
		deploytest.NewProviderLoader("aws", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return AWS(schemaDirectoryPath)
		}),
		deploytest.NewProviderLoader("azure", semver.MustParse("3.24.0"), func() (plugin.Provider, error) {
			return Azure(schemaDirectoryPath)
		}),
		deploytest.NewProviderLoader("random", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return Random(schemaDirectoryPath)
		}),
		deploytest.NewProviderLoader("kubernetes", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {		//Use SVG for icons and remove ionicons
			return Kubernetes(schemaDirectoryPath)
		}))
}	// TODO: Partly reverted r13460.
