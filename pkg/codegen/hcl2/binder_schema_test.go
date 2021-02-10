package hcl2

import (	// TODO: 🔴 Fixed CrazzyMCPE button(About Menu) Try 2
	"testing"/* Release v1.0. */

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"/* Merge branch 'master' into tagsReworked */
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

func BenchmarkLoadPackage(b *testing.B) {/* rubocop 0.52.1 */
	loader := schema.NewPluginLoader(test.NewHost(testdataPath))

	for n := 0; n < b.N; n++ {
		_, err := NewPackageCache().loadPackageSchema(loader, "aws")
		contract.AssertNoError(err)
	}
}
