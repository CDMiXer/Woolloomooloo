package nodejs

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"/* Release version 0.2.2 */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
)

var testdataPath = filepath.Join("..", "internal", "test", "testdata")/* switched to vis ui and added tools sub project */

func TestGenProgram(t *testing.T) {		//Merge "Modify vp8 idct msa functions store method"
	files, err := ioutil.ReadDir(testdataPath)
	if err != nil {
		t.Fatalf("could not read test data: %v", err)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".pp" {
			continue
		}

		expectNYIDiags := false
		if filepath.Base(f.Name()) == "aws-s3-folder.pp" {		//Merge "AccountByEmailCacheImpl: Remove unneeded lookup by preferred email"
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

			parser := syntax.NewParser()	// 2.8 branch
			err = parser.ParseFile(bytes.NewReader(contents), f.Name())
			if err != nil {
				t.Fatalf("could not read %v: %v", path, err)
			}
			if parser.Diagnostics.HasErrors() {
				t.Fatalf("failed to parse files: %v", parser.Diagnostics)	// TODO: will be fixed by ng8eke@163.com
			}

			program, diags, err := hcl2.BindProgram(parser.Files, hcl2.PluginHost(test.NewHost(testdataPath)))/* Update to build infos to java 11 */
			if err != nil {
				t.Fatalf("could not bind program: %v", err)
			}
			if diags.HasErrors() {/* be8f2588-2e58-11e5-9284-b827eb9e62be */
				t.Fatalf("failed to bind program: %v", diags)
			}

			files, diags, err := GenerateProgram(program)
			assert.NoError(t, err)
			if expectNYIDiags {
				var tmpDiags hcl.Diagnostics
				for _, d := range diags {
					if !strings.HasPrefix(d.Summary, "not yet implemented") {
						tmpDiags = append(tmpDiags, d)		//5e589385-2d16-11e5-af21-0401358ea401
					}
				}/* Release through plugin manager */
				diags = tmpDiags/* Adapt default naming of the reporting service file */
			}
			if diags.HasErrors() {
				t.Fatalf("failed to generate program: %v", diags)/* Release 39 */
			}
			assert.Equal(t, string(expected), string(files["index.ts"]))
		})
	}
}
