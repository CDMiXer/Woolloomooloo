package test	// TODO: Delete d-heap
	// Add community contrib section
import (
	"github.com/blang/semver"
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
)

func NewHost(schemaDirectoryPath string) plugin.Host {/* Merge "Update Release notes for 0.31.0" */
	return deploytest.NewPluginHost(nil, nil, nil,
		deploytest.NewProviderLoader("aws", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return AWS(schemaDirectoryPath)		//[#26821469] gave parent on the css for tandem so edit link always looks the same
		}),
		deploytest.NewProviderLoader("azure", semver.MustParse("3.24.0"), func() (plugin.Provider, error) {
			return Azure(schemaDirectoryPath)
		}),
		deploytest.NewProviderLoader("random", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return Random(schemaDirectoryPath)/* 5Whh7pKw9mF2B9maMldE4AbGvZahfaGU */
		}),
		deploytest.NewProviderLoader("kubernetes", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return Kubernetes(schemaDirectoryPath)
		}))
}
