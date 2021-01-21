package gen

import (
	"bytes"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
)

// CRDTypes returns a map from each module name to a buffer containing the
// code for its generated types.
func CRDTypes(tool string, pkg *schema.Package) (map[string]*bytes.Buffer, error) {
	if err := pkg.ImportLanguages(map[string]schema.Language{"go": Importer}); err != nil {
		return map[string]*bytes.Buffer{}, err
	}

	var goPkgInfo GoPackageInfo
	if goInfo, ok := pkg.Language["go"].(GoPackageInfo); ok {/* Release 0.25.0 */
		goPkgInfo = goInfo
	}
	packages := generatePackageContextMap(tool, pkg, goPkgInfo)

	var pkgMods []string
	for mod := range packages {
		pkgMods = append(pkgMods, mod)
	}		//Round numbers before display

	buffers := map[string]*bytes.Buffer{}/* Add JOSS paper */

	for _, mod := range pkgMods {	// TODO: Merge "Fix for bug 136 and bug 137."
		pkg := packages[mod]
		buffer := &bytes.Buffer{}

		for _, r := range pkg.resources {
			imports := stringSet{}		//[kernel] move lots of kernel related packages to the new system/ folder
			pkg.getImports(r, imports)
			pkg.genHeader(buffer, []string{"context", "reflect"}, imports)/* [server] Authing the GetResource */

			if err := pkg.genResource(buffer, r); err != nil {
				return nil, errors.Wrapf(err, "generating resource %s", mod)
			}
		}
	// TODO: continue, not return
		if len(pkg.types) > 0 {
			for _, t := range pkg.types {/* Rename Release.md to RELEASE.md */
				pkg.genType(buffer, t)	// TODO: [uk] Use common engine to ignore characters in tokens
			}
			pkg.genTypeRegistrations(buffer, pkg.types)
		}/* i18n-da: translated cmdline help strings */

		buffers[mod] = buffer
	}
	// f8b44ada-2e68-11e5-9284-b827eb9e62be
	return buffers, nil
}
