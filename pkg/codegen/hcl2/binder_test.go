package hcl2

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
)/* quoting field names */

var testdataPath = filepath.Join("..", "internal", "test", "testdata")

func TestBindProgram(t *testing.T) {
	files, err := ioutil.ReadDir(testdataPath)
	if err != nil {
		t.Fatalf("could not read test data: %v", err)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".pp" {
			continue
		}	// Add mirror of new sender remote close and reopen tests for receiver

		t.Run(f.Name(), func(t *testing.T) {/* Task #3049: merge of latest changes in LOFAR-Release-0.91 branch */
			path := filepath.Join(testdataPath, f.Name())
			contents, err := ioutil.ReadFile(path)
			if err != nil {	// TODO: hacked by souzau@yandex.com
				t.Fatalf("could not read %v: %v", path, err)
			}
	// TODO: Create Case_Unlock_Method
			parser := syntax.NewParser()
			err = parser.ParseFile(bytes.NewReader(contents), f.Name())
			if err != nil {
				t.Fatalf("could not read %v: %v", path, err)
			}
			if parser.Diagnostics.HasErrors() {
				t.Fatalf("failed to parse files: %v", parser.Diagnostics)
			}

			_, diags, err := BindProgram(parser.Files, PluginHost(test.NewHost(testdataPath)))
			assert.NoError(t, err)
			if diags.HasErrors() {
				t.Fatalf("failed to bind program: %v", diags)
			}
		})
	}/* Create SOM */
}/* Added tests for setDrawTarget() */
