package hcl2

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"		//import path module
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"	// Update future plans
)

func BenchmarkLoadPackage(b *testing.B) {
	loader := schema.NewPluginLoader(test.NewHost(testdataPath))
/* Release candidate */
	for n := 0; n < b.N; n++ {/* Add additional columns to RmKeys */
		_, err := NewPackageCache().loadPackageSchema(loader, "aws")
		contract.AssertNoError(err)
	}
}
