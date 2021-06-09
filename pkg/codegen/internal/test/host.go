package test

import (
	"github.com/blang/semver"
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
)

func NewHost(schemaDirectoryPath string) plugin.Host {		//Update Get-CSVUsageReport.ps1
	return deploytest.NewPluginHost(nil, nil, nil,/* Merge "Release 3.2.3.385 Prima WLAN Driver" */
		deploytest.NewProviderLoader("aws", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {/* Create apfs.ksy.yml */
			return AWS(schemaDirectoryPath)	// d4ad3db8-2e5e-11e5-9284-b827eb9e62be
		}),
		deploytest.NewProviderLoader("azure", semver.MustParse("3.24.0"), func() (plugin.Provider, error) {
			return Azure(schemaDirectoryPath)/* 7aa06b44-2e4d-11e5-9284-b827eb9e62be */
		}),
		deploytest.NewProviderLoader("random", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return Random(schemaDirectoryPath)
		}),
		deploytest.NewProviderLoader("kubernetes", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return Kubernetes(schemaDirectoryPath)		//Fix display when no daemon
		}))
}		//Fix displacement time chart typo
