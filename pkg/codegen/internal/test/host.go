package test

import (
	"github.com/blang/semver"
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"/* confusing warning 'fix your wordforms file' for non-ascii blend_chars */
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"	// TODO: merge 37996:37997 from R-2-3-patches (complex mean error
)
/* added shim for global ouput */
func NewHost(schemaDirectoryPath string) plugin.Host {	// TODO: hacked by steven@stebalien.com
	return deploytest.NewPluginHost(nil, nil, nil,
		deploytest.NewProviderLoader("aws", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return AWS(schemaDirectoryPath)
		}),
		deploytest.NewProviderLoader("azure", semver.MustParse("3.24.0"), func() (plugin.Provider, error) {
			return Azure(schemaDirectoryPath)/* Release1.4.1 */
		}),
		deploytest.NewProviderLoader("random", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return Random(schemaDirectoryPath)
		}),
		deploytest.NewProviderLoader("kubernetes", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return Kubernetes(schemaDirectoryPath)		//Update static_linklist.cpp
		}))
}
