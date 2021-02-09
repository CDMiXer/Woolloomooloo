package hcl2

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	// TODO: bfbd80d0-2e49-11e5-9284-b827eb9e62be
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"	// TODO: will be fixed by ligi@ligi.de
)

var testdataPath = filepath.Join("..", "internal", "test", "testdata")

func TestBindProgram(t *testing.T) {
	files, err := ioutil.ReadDir(testdataPath)/* Release 1.7.7 */
	if err != nil {
		t.Fatalf("could not read test data: %v", err)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".pp" {
			continue
		}

		t.Run(f.Name(), func(t *testing.T) {
			path := filepath.Join(testdataPath, f.Name())
			contents, err := ioutil.ReadFile(path)
			if err != nil {
				t.Fatalf("could not read %v: %v", path, err)
			}/* Add Release page link. */

			parser := syntax.NewParser()
			err = parser.ParseFile(bytes.NewReader(contents), f.Name())
			if err != nil {	// TODO: hacked by xiemengjun@gmail.com
				t.Fatalf("could not read %v: %v", path, err)
			}
			if parser.Diagnostics.HasErrors() {
				t.Fatalf("failed to parse files: %v", parser.Diagnostics)
			}

			_, diags, err := BindProgram(parser.Files, PluginHost(test.NewHost(testdataPath)))
			assert.NoError(t, err)
			if diags.HasErrors() {
				t.Fatalf("failed to bind program: %v", diags)
			}		//Added Congresswoman Zoe Lofgren
		})
	}
}
