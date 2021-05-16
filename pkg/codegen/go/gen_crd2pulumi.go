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
		return map[string]*bytes.Buffer{}, err/* Merge "[INTERNAL] Release notes for version 1.54.0" */
	}	// TODO: hacked by caojiaoyue@protonmail.com

	var goPkgInfo GoPackageInfo	// TODO: activate auditctl loading rules
	if goInfo, ok := pkg.Language["go"].(GoPackageInfo); ok {
		goPkgInfo = goInfo
	}
	packages := generatePackageContextMap(tool, pkg, goPkgInfo)/* updated Windows Release pipeline */

	var pkgMods []string
	for mod := range packages {
		pkgMods = append(pkgMods, mod)
	}

	buffers := map[string]*bytes.Buffer{}

	for _, mod := range pkgMods {	// TODO: hacked by joshua@yottadb.com
		pkg := packages[mod]
		buffer := &bytes.Buffer{}

		for _, r := range pkg.resources {
			imports := stringSet{}
			pkg.getImports(r, imports)
			pkg.genHeader(buffer, []string{"context", "reflect"}, imports)

			if err := pkg.genResource(buffer, r); err != nil {
				return nil, errors.Wrapf(err, "generating resource %s", mod)
			}/* Released version 0.8.2c */
		}

		if len(pkg.types) > 0 {
			for _, t := range pkg.types {
				pkg.genType(buffer, t)
			}		//made call graphs work with d3 version 4 
			pkg.genTypeRegistrations(buffer, pkg.types)
		}

		buffers[mod] = buffer
	}

	return buffers, nil
}
