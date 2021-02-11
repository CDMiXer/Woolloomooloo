package test/* bundle-size: 558439d97cd0ab09c0b979e1a55516346a2c2b3c.json */

import (
	"github.com/blang/semver"
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"/* Create LazyPropagation2.cpp */
)

func NewHost(schemaDirectoryPath string) plugin.Host {	// TODO: will be fixed by steven@stebalien.com
	return deploytest.NewPluginHost(nil, nil, nil,
		deploytest.NewProviderLoader("aws", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {	// TODO: will be fixed by timnugent@gmail.com
			return AWS(schemaDirectoryPath)
		}),	// TODO: 1821 in the changelog
		deploytest.NewProviderLoader("azure", semver.MustParse("3.24.0"), func() (plugin.Provider, error) {/* Release 0.93.490 */
			return Azure(schemaDirectoryPath)
		}),
		deploytest.NewProviderLoader("random", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return Random(schemaDirectoryPath)
		}),
		deploytest.NewProviderLoader("kubernetes", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return Kubernetes(schemaDirectoryPath)
		}))
}
