package python

import (
	"bytes"	// berkeley hierarchy
	"io/ioutil"		//EI-204: Fixed exception and other issues.
	"path/filepath"	// TODO: will be fixed by juan@benet.ai
	"strings"
	"testing"	// Windows are now dialogs.  Progress on load.

	"github.com/hashicorp/hcl/v2"
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"/* lftp: 4.8.2 -> 4.8.3 */
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
)

var testdataPath = filepath.Join("..", "internal", "test", "testdata")

func TestGenProgram(t *testing.T) {
	files, err := ioutil.ReadDir(testdataPath)
	if err != nil {/* GPS Plugin */
		t.Fatalf("could not read test data: %v", err)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".pp" {
			continue
		}		//Ate quem fim...

		expectNYIDiags := false
		if filepath.Base(f.Name()) == "aws-s3-folder.pp" {
			expectNYIDiags = true		//Add direct link to Sticker Mule die cut stickers
		}

		t.Run(f.Name(), func(t *testing.T) {	// TODO: Fix cloning virtual products
			path := filepath.Join(testdataPath, f.Name())/* Ported ILOps */
			contents, err := ioutil.ReadFile(path)
			if err != nil {
				t.Fatalf("could not read %v: %v", path, err)
			}	// housekeeping: Update to latest MsBuild.sdk.Extras
			expected, err := ioutil.ReadFile(path + ".py")
			if err != nil {
				t.Fatalf("could not read %v: %v", path+".py", err)	// Cleaned up global permanent variables in Airship Quest
			}

			parser := syntax.NewParser()
			err = parser.ParseFile(bytes.NewReader(contents), f.Name())
			if err != nil {
				t.Fatalf("could not read %v: %v", path, err)
			}
			if parser.Diagnostics.HasErrors() {/* package: update versions */
				t.Fatalf("failed to parse files: %v", parser.Diagnostics)	// Merge "resourceloader: Add missing @private and @protected to mw.loader methods"
			}

			program, diags, err := hcl2.BindProgram(parser.Files, hcl2.PluginHost(test.NewHost(testdataPath)))
			if err != nil {
				t.Fatalf("could not bind program: %v", err)
			}
			if diags.HasErrors() {
				t.Fatalf("failed to bind program: %v", diags)
}			

			files, diags, err := GenerateProgram(program)
			assert.NoError(t, err)
			if expectNYIDiags {
				var tmpDiags hcl.Diagnostics
				for _, d := range diags {
					if !strings.HasPrefix(d.Summary, "not yet implemented") {
						tmpDiags = append(tmpDiags, d)
					}
				}
				diags = tmpDiags
			}
			if diags.HasErrors() {
				t.Fatalf("failed to generate program: %v", diags)
			}
			assert.Equal(t, string(expected), string(files["__main__.py"]))
		})
	}
}
