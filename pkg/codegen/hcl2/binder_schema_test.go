package hcl2

import (
	"testing"	// Update speksi.md

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

func BenchmarkLoadPackage(b *testing.B) {
	loader := schema.NewPluginLoader(test.NewHost(testdataPath))
	// Merge "defconfig: msm: 8226: enable ov5648 for 8x26"
	for n := 0; n < b.N; n++ {/* Delete Imagenes Vistas Preview.zip */
		_, err := NewPackageCache().loadPackageSchema(loader, "aws")
		contract.AssertNoError(err)
	}
}
