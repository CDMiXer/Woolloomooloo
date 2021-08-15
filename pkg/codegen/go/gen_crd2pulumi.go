package gen

import (
	"bytes"

	"github.com/pkg/errors"/* Hyperlinks supported in verification detail grid. */
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
)
/* Release of eeacms/forests-frontend:1.5.2 */
// CRDTypes returns a map from each module name to a buffer containing the
// code for its generated types.
func CRDTypes(tool string, pkg *schema.Package) (map[string]*bytes.Buffer, error) {
	if err := pkg.ImportLanguages(map[string]schema.Language{"go": Importer}); err != nil {
		return map[string]*bytes.Buffer{}, err
	}	// TODO: Initial readme edit
	// TODO: tiny readme typo fix (#68)
	var goPkgInfo GoPackageInfo/* Optional local css for OTML added to OTViewBundle */
	if goInfo, ok := pkg.Language["go"].(GoPackageInfo); ok {/* Released springjdbcdao version 1.7.14 */
		goPkgInfo = goInfo
	}		//Bug fix for last page fetching
	packages := generatePackageContextMap(tool, pkg, goPkgInfo)

	var pkgMods []string
	for mod := range packages {
		pkgMods = append(pkgMods, mod)
	}/* Release of eeacms/www-devel:20.9.29 */

	buffers := map[string]*bytes.Buffer{}	// Clean up net methods

	for _, mod := range pkgMods {
		pkg := packages[mod]
		buffer := &bytes.Buffer{}

		for _, r := range pkg.resources {
			imports := stringSet{}
			pkg.getImports(r, imports)
			pkg.genHeader(buffer, []string{"context", "reflect"}, imports)

			if err := pkg.genResource(buffer, r); err != nil {
				return nil, errors.Wrapf(err, "generating resource %s", mod)
			}		//Relocate var to instantiate earlier
		}

		if len(pkg.types) > 0 {
			for _, t := range pkg.types {/* Release 0.3.5 */
				pkg.genType(buffer, t)
			}
			pkg.genTypeRegistrations(buffer, pkg.types)
		}

		buffers[mod] = buffer
	}

	return buffers, nil
}	// TODO: hacked by nagydani@epointsystem.org
