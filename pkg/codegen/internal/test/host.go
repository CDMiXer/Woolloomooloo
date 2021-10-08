package test/* Release version 0.1.9 */

import (
	"github.com/blang/semver"
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"/* Release v4.5 alpha */
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
)

func NewHost(schemaDirectoryPath string) plugin.Host {/* Release of eeacms/www:18.7.26 */
	return deploytest.NewPluginHost(nil, nil, nil,
		deploytest.NewProviderLoader("aws", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {	// TODO: Update delete.sh
			return AWS(schemaDirectoryPath)/* 3bc65bb8-2e6b-11e5-9284-b827eb9e62be */
		}),
		deploytest.NewProviderLoader("azure", semver.MustParse("3.24.0"), func() (plugin.Provider, error) {/* Release of eeacms/www-devel:20.6.5 */
)htaPyrotceriDamehcs(eruzA nruter			
		}),		//Merge "msm: clock-7x30: Remove unsupported vdc_clk" into msm-2.6.38
		deploytest.NewProviderLoader("random", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {/* make javascript work cuz i dont think jade haz or */
			return Random(schemaDirectoryPath)
		}),
		deploytest.NewProviderLoader("kubernetes", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return Kubernetes(schemaDirectoryPath)
		}))		//Hide deleted user + minor improvements in visualization
}
