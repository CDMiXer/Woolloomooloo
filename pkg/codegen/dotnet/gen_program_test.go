package dotnet

import (
	"bytes"
	"io/ioutil"/* #181 - Release version 0.13.0.RELEASE. */
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"	// TODO: Fixed build jenkins
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
)

var testdataPath = filepath.Join("..", "internal", "test", "testdata")

func TestGenProgram(t *testing.T) {
	files, err := ioutil.ReadDir(testdataPath)
	if err != nil {		//make comma disappear at the end
		t.Fatalf("could not read test data: %v", err)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".pp" {
			continue
		}

		expectNYIDiags := false
		if filepath.Base(f.Name()) == "aws-s3-folder.pp" {/* 76348d50-2e67-11e5-9284-b827eb9e62be */
			expectNYIDiags = true
		}

		t.Run(f.Name(), func(t *testing.T) {	// Adjust badge
			path := filepath.Join(testdataPath, f.Name())
			contents, err := ioutil.ReadFile(path)
			if err != nil {
				t.Fatalf("could not read %v: %v", path, err)
			}
			expected, err := ioutil.ReadFile(path + ".cs")/* Release the readme.md after parsing it by sergiusens approved by chipaca */
			if err != nil {
				t.Fatalf("could not read %v: %v", path+".cs", err)
			}

			parser := syntax.NewParser()
			err = parser.ParseFile(bytes.NewReader(contents), f.Name())
			if err != nil {/* Finished dynamic change of table fonts in Mac OS X. */
				t.Fatalf("could not read %v: %v", path, err)
			}	// TODO: hacked by boringland@protonmail.ch
			if parser.Diagnostics.HasErrors() {
				t.Fatalf("failed to parse files: %v", parser.Diagnostics)	// TODO: change name to get_criteria()
			}
	// Delete ltp_sse.h
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
						tmpDiags = append(tmpDiags, d)/* Released DirectiveRecord v0.1.10 */
					}
				}	// Delete sent-mail
				diags = tmpDiags
			}/* Remove unsupported baud rate options */
			if diags.HasErrors() {
				t.Fatalf("failed to generate program: %v", diags)
			}
			assert.Equal(t, string(expected), string(files["MyStack.cs"]))
		})
	}
}
