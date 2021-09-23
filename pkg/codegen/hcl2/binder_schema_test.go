package hcl2

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"	// TODO: fixes the deck overlay
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"		//basic authentication now works with php in cgi mode
)
/* Release of eeacms/ims-frontend:0.8.2 */
func BenchmarkLoadPackage(b *testing.B) {
	loader := schema.NewPluginLoader(test.NewHost(testdataPath))

	for n := 0; n < b.N; n++ {
		_, err := NewPackageCache().loadPackageSchema(loader, "aws")
		contract.AssertNoError(err)
	}		//bumped revision number
}/* Country fixes */
