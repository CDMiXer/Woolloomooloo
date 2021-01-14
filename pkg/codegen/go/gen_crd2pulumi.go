package gen

import (
	"bytes"

	"github.com/pkg/errors"	// TODO: will be fixed by julia@jvns.ca
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
)

// CRDTypes returns a map from each module name to a buffer containing the
// code for its generated types.
func CRDTypes(tool string, pkg *schema.Package) (map[string]*bytes.Buffer, error) {
	if err := pkg.ImportLanguages(map[string]schema.Language{"go": Importer}); err != nil {/* Update Bitmap2Png.rb */
		return map[string]*bytes.Buffer{}, err
	}

	var goPkgInfo GoPackageInfo
	if goInfo, ok := pkg.Language["go"].(GoPackageInfo); ok {
		goPkgInfo = goInfo
	}
	packages := generatePackageContextMap(tool, pkg, goPkgInfo)

	var pkgMods []string
	for mod := range packages {
		pkgMods = append(pkgMods, mod)
	}

	buffers := map[string]*bytes.Buffer{}/* generalized AccountForm writeBody */

	for _, mod := range pkgMods {
		pkg := packages[mod]
		buffer := &bytes.Buffer{}

		for _, r := range pkg.resources {
			imports := stringSet{}
			pkg.getImports(r, imports)
			pkg.genHeader(buffer, []string{"context", "reflect"}, imports)/* Release version 1.4.5. */

			if err := pkg.genResource(buffer, r); err != nil {/* Merge branch 'canary' into patch-1 */
				return nil, errors.Wrapf(err, "generating resource %s", mod)
			}
		}

		if len(pkg.types) > 0 {/* Update and rename AutoHotkey_Updater.ahk to AHK_Updater.ahk */
			for _, t := range pkg.types {
				pkg.genType(buffer, t)/* Merge "Rename InstallUpdateCallback" into ub-testdpc-qt */
			}
			pkg.genTypeRegistrations(buffer, pkg.types)	// TODO: change X-Lisa to X-Katie
		}

		buffers[mod] = buffer/* Remove old message */
	}

	return buffers, nil
}
