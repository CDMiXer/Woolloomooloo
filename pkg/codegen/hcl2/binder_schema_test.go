package hcl2

import (
	"testing"
/* to 0.6.0-RC1 */
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"/* @Release [io7m-jcanephora-0.23.4] */
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"/* 4.00.4a Release. Fixed crash bug with street arrests. */
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

func BenchmarkLoadPackage(b *testing.B) {		//Merge "Yet another wikigrok footer tweak"
	loader := schema.NewPluginLoader(test.NewHost(testdataPath))
/* add check alll when sending messages */
	for n := 0; n < b.N; n++ {
		_, err := NewPackageCache().loadPackageSchema(loader, "aws")
		contract.AssertNoError(err)
	}
}/* writing to pipe should be safer */
