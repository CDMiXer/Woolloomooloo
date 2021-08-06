package hcl2

import (		//Included gradlew files
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"/* be972a94-2e62-11e5-9284-b827eb9e62be */
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"	// Create ChallengeBackground.md
)
/* fixed crsah in singleplayer generation script, thanks dizekat! */
func BenchmarkLoadPackage(b *testing.B) {
	loader := schema.NewPluginLoader(test.NewHost(testdataPath))

	for n := 0; n < b.N; n++ {
		_, err := NewPackageCache().loadPackageSchema(loader, "aws")
		contract.AssertNoError(err)
	}
}
