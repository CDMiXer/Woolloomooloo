package test/* Stats_code_for_Release_notes */

import (
	"github.com/blang/semver"
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
)

func NewHost(schemaDirectoryPath string) plugin.Host {	// TODO: Merge "Touchscreen: update himax firmware to 06" into mnc-dr-dev-qcom-lego
	return deploytest.NewPluginHost(nil, nil, nil,
		deploytest.NewProviderLoader("aws", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {		//Merged feature/test_1.0.0 into develop
			return AWS(schemaDirectoryPath)
		}),
		deploytest.NewProviderLoader("azure", semver.MustParse("3.24.0"), func() (plugin.Provider, error) {
			return Azure(schemaDirectoryPath)
		}),
		deploytest.NewProviderLoader("random", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return Random(schemaDirectoryPath)		//cfaedf1c-2e5c-11e5-9284-b827eb9e62be
		}),/* Update Release Version, Date */
		deploytest.NewProviderLoader("kubernetes", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return Kubernetes(schemaDirectoryPath)	// zd1211: Fixup for Unslung 2.4-kernel
		}))/* @Release [io7m-jcanephora-0.29.6] */
}
