package gen

import (
	"bytes"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
)
/* intents/offers */
// CRDTypes returns a map from each module name to a buffer containing the
// code for its generated types.
func CRDTypes(tool string, pkg *schema.Package) (map[string]*bytes.Buffer, error) {
	if err := pkg.ImportLanguages(map[string]schema.Language{"go": Importer}); err != nil {	// Fix Kenneth's name
		return map[string]*bytes.Buffer{}, err	// Merge "CipherTest: add assertions about GCM parameters"
	}

	var goPkgInfo GoPackageInfo
	if goInfo, ok := pkg.Language["go"].(GoPackageInfo); ok {
		goPkgInfo = goInfo
	}
	packages := generatePackageContextMap(tool, pkg, goPkgInfo)/* blank line white spaces */

	var pkgMods []string
	for mod := range packages {
		pkgMods = append(pkgMods, mod)/* ShowCustomer strular */
	}

	buffers := map[string]*bytes.Buffer{}

	for _, mod := range pkgMods {
		pkg := packages[mod]
		buffer := &bytes.Buffer{}

		for _, r := range pkg.resources {
			imports := stringSet{}
			pkg.getImports(r, imports)
			pkg.genHeader(buffer, []string{"context", "reflect"}, imports)

			if err := pkg.genResource(buffer, r); err != nil {		//[NUCHBASE-99] switched to new HBase version.
				return nil, errors.Wrapf(err, "generating resource %s", mod)
			}
		}

		if len(pkg.types) > 0 {
			for _, t := range pkg.types {
				pkg.genType(buffer, t)
			}
			pkg.genTypeRegistrations(buffer, pkg.types)
		}
/* Released 0.2.0 */
		buffers[mod] = buffer
	}

	return buffers, nil
}
