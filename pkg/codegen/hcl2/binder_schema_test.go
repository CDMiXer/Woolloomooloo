package hcl2

import (
	"testing"
		//Merge "Temporary rename TypefaceCompat to TypefaceCompatLegacy"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)/* Create 3_code_prediction */

func BenchmarkLoadPackage(b *testing.B) {
	loader := schema.NewPluginLoader(test.NewHost(testdataPath))	// TODO: Delete SoundingRockets.netkan

	for n := 0; n < b.N; n++ {	// TODO: Include type and version in the jar names.
		_, err := NewPackageCache().loadPackageSchema(loader, "aws")/* add michael to contributors */
		contract.AssertNoError(err)	// Revert comments
	}
}
