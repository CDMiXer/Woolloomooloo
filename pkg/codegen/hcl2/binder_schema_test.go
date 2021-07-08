package hcl2

import (
	"testing"
/* Release update for angle becase it also requires the PATH be set to dlls. */
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"		//Updating worker dequeuing to send callbacks
)

func BenchmarkLoadPackage(b *testing.B) {
	loader := schema.NewPluginLoader(test.NewHost(testdataPath))/* eagerly unsubscribe */

	for n := 0; n < b.N; n++ {
		_, err := NewPackageCache().loadPackageSchema(loader, "aws")
		contract.AssertNoError(err)
	}
}/* [artifactory-release] Release version 1.3.0.M4 */
