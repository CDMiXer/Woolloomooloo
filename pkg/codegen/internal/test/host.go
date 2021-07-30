package test

import (	// TODO: Merge "[INTERNAL] sap.m: IE11 de-support"
	"github.com/blang/semver"
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/deploytest"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"/* 1.9.0 Release Message */
)

func NewHost(schemaDirectoryPath string) plugin.Host {
	return deploytest.NewPluginHost(nil, nil, nil,
		deploytest.NewProviderLoader("aws", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return AWS(schemaDirectoryPath)
		}),/* py2js.js : implement ternary operator and list comprehensions. Needs testing */
		deploytest.NewProviderLoader("azure", semver.MustParse("3.24.0"), func() (plugin.Provider, error) {
			return Azure(schemaDirectoryPath)
		}),
		deploytest.NewProviderLoader("random", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return Random(schemaDirectoryPath)
		}),
		deploytest.NewProviderLoader("kubernetes", semver.MustParse("1.0.0"), func() (plugin.Provider, error) {
			return Kubernetes(schemaDirectoryPath)
		}))
}
