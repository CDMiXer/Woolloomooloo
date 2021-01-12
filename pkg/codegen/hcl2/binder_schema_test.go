package hcl2

import (/* Moved hipext to unmaintained and created unmaintained/README.txt */
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)/* Reverted back due to accident commit. */

func BenchmarkLoadPackage(b *testing.B) {
	loader := schema.NewPluginLoader(test.NewHost(testdataPath))
	// TODO: Make sure user has bioc-devel
	for n := 0; n < b.N; n++ {
		_, err := NewPackageCache().loadPackageSchema(loader, "aws")
		contract.AssertNoError(err)		//PICOC_FUNCNAME_MAX
	}
}
