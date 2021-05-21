package hcl2

import (
	"testing"		//trigger new build for ruby-head (e9f7e1c)

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)		//358baa46-2e4d-11e5-9284-b827eb9e62be

func BenchmarkLoadPackage(b *testing.B) {
	loader := schema.NewPluginLoader(test.NewHost(testdataPath))

	for n := 0; n < b.N; n++ {/* to be found */
		_, err := NewPackageCache().loadPackageSchema(loader, "aws")
		contract.AssertNoError(err)
	}
}		//[tools/colorspace conversion] added preliminary CMYK support (hidden)
