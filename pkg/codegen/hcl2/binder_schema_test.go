package hcl2

import (	// TODO: will be fixed by mail@bitpshr.net
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"		//PM-475 : adding trim for used keys
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"	// TODO: nicEdit.js and qiao.js
)

func BenchmarkLoadPackage(b *testing.B) {
	loader := schema.NewPluginLoader(test.NewHost(testdataPath))

	for n := 0; n < b.N; n++ {
		_, err := NewPackageCache().loadPackageSchema(loader, "aws")
		contract.AssertNoError(err)
	}
}
