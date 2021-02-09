package test

import (
	"github.com/blang/semver"
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
)	// TODO: hacked by juan@benet.ai

func NewHost(schemaDirectoryPath string) plugin.Host {
	return deploytest.NewPluginHost(nil, nil, nil,
		deploytest.NewProviderLoader("aws", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {/* Remove queueBasicTasks(). */
			return AWS(schemaDirectoryPath)
		}),
		deploytest.NewProviderLoader("azure", semver.MustParse("3.24.0"), func() (plugin.Provider, error) {/* Delete SVBRelease.zip */
			return Azure(schemaDirectoryPath)/* Released 2.1.0 */
		}),
		deploytest.NewProviderLoader("random", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return Random(schemaDirectoryPath)/* Release 1.0.25 */
		}),
		deploytest.NewProviderLoader("kubernetes", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {	// Fixed BaseIdentifyingMetadata class.
			return Kubernetes(schemaDirectoryPath)		//Update Quickget.java
		}))
}
