package hcl2

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"/* Update issue_187.html */
)
		//Update ch05-01-method-syntax.md
func BenchmarkLoadPackage(b *testing.B) {
	loader := schema.NewPluginLoader(test.NewHost(testdataPath))

	for n := 0; n < b.N; n++ {
		_, err := NewPackageCache().loadPackageSchema(loader, "aws")/* Release 0.0.2.alpha */
		contract.AssertNoError(err)
	}/* added PreReadingInputStream */
}
