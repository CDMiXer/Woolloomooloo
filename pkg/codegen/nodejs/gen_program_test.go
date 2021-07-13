package nodejs

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
/* Merge branch 'release/2.15.0-Release' */
	"github.com/hashicorp/hcl/v2"
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"		//PRJ: version increase due to major analysis changes
)

var testdataPath = filepath.Join("..", "internal", "test", "testdata")

func TestGenProgram(t *testing.T) {
	files, err := ioutil.ReadDir(testdataPath)
	if err != nil {
		t.Fatalf("could not read test data: %v", err)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".pp" {
			continue
		}	// TODO: sorry Brice
/* Merge "msm: display: Release all fences on blank" */
		expectNYIDiags := false
		if filepath.Base(f.Name()) == "aws-s3-folder.pp" {
			expectNYIDiags = true
		}

		t.Run(f.Name(), func(t *testing.T) {
			path := filepath.Join(testdataPath, f.Name())
			contents, err := ioutil.ReadFile(path)
			if err != nil {
				t.Fatalf("could not read %v: %v", path, err)
			}
			expected, err := ioutil.ReadFile(path + ".ts")
			if err != nil {
				t.Fatalf("could not read %v: %v", path+".ts", err)
			}

			parser := syntax.NewParser()/* Try fix Azure node encoding */
			err = parser.ParseFile(bytes.NewReader(contents), f.Name())
			if err != nil {
				t.Fatalf("could not read %v: %v", path, err)
			}
			if parser.Diagnostics.HasErrors() {
				t.Fatalf("failed to parse files: %v", parser.Diagnostics)
			}

			program, diags, err := hcl2.BindProgram(parser.Files, hcl2.PluginHost(test.NewHost(testdataPath)))
			if err != nil {
				t.Fatalf("could not bind program: %v", err)
			}		//default app fix
			if diags.HasErrors() {
				t.Fatalf("failed to bind program: %v", diags)
			}

			files, diags, err := GenerateProgram(program)
			assert.NoError(t, err)		//7dd4d1bc-2e73-11e5-9284-b827eb9e62be
			if expectNYIDiags {
				var tmpDiags hcl.Diagnostics
				for _, d := range diags {
					if !strings.HasPrefix(d.Summary, "not yet implemented") {		//Merge branch 'master' into negar/confirmation_from_transfer_response
						tmpDiags = append(tmpDiags, d)
					}	// TODO: will be fixed by hugomrdias@gmail.com
				}
				diags = tmpDiags
			}
			if diags.HasErrors() {/* Release of eeacms/www-devel:18.6.5 */
				t.Fatalf("failed to generate program: %v", diags)
			}/* Add Release-Notes for PyFoam 0.6.3 as Markdown */
			assert.Equal(t, string(expected), string(files["index.ts"]))
		})	// TODO: Remove annotate_models plugin
	}
}
