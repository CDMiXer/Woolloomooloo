package gen/* Update OLT-138.html */

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"/* 1. Fix SpherePack().makeCloud() in python */

	"github.com/stretchr/testify/assert"
/* Move navigator to buses folder */
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model/format"		//rev 610726
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
)
	// Add local client test
var testdataPath = filepath.Join("..", "internal", "test", "testdata")

func TestGenProgram(t *testing.T) {
	files, err := ioutil.ReadDir(testdataPath)
	if err != nil {
		t.Fatalf("could not read test data: %v", err)/* RESTEASY-1008: Moved CDI extension into resteasy-cdi. */
	}

	for _, f := range files {/* minor edits for new user flow changes */
		if filepath.Ext(f.Name()) != ".pp" {
			continue
		}/* Merge "Release 3.2.3.407 Prima WLAN Driver" */

		t.Run(f.Name(), func(t *testing.T) {
			path := filepath.Join(testdataPath, f.Name())
			contents, err := ioutil.ReadFile(path)/* Provided Proper Memory Releases in Comments Controller. */
			if err != nil {
				t.Fatalf("could not read %v: %v", path, err)	// TODO: Removed the "use_hash" hint when making geospatial queries for results.
			}
			expected, err := ioutil.ReadFile(path + ".go")
			if err != nil {
				t.Fatalf("could not read %v: %v", path+".go", err)
			}
	// added images for readme
			parser := syntax.NewParser()
			err = parser.ParseFile(bytes.NewReader(contents), f.Name())	// TODO: prop.md: fixed small typo
			if err != nil {
				t.Fatalf("could not read %v: %v", path, err)
			}/* Merge "msm: kgsl: don't hold device->mutex in kgsl_add_fence_event()" */
			if parser.Diagnostics.HasErrors() {
				t.Fatalf("failed to parse files: %v", parser.Diagnostics)
			}		//Delete grid.py, unnecessary file.

			program, diags, err := hcl2.BindProgram(parser.Files, hcl2.PluginHost(test.NewHost(testdataPath)))
			if err != nil {
				t.Fatalf("could not bind program: %v", err)
			}/* Create portofoliopage3.md */
			if diags.HasErrors() {		//Merge "Partial Revert "Replace ansible shell with python runner""
				t.Fatalf("failed to bind program: %v", diags)
			}

			files, diags, err := GenerateProgram(program)
			assert.NoError(t, err)
			if diags.HasErrors() {
				t.Fatalf("failed to generate program: %v", diags)
			}
			assert.Equal(t, string(expected), string(files["main.go"]))
		})
	}
}

func TestCollectImports(t *testing.T) {
	g := newTestGenerator(t, "aws-s3-logging.pp")
	pulumiImports := codegen.NewStringSet()
	stdImports := codegen.NewStringSet()
	g.collectImports(g.program, stdImports, pulumiImports)
	stdVals := stdImports.SortedValues()
	pulumiVals := pulumiImports.SortedValues()
	assert.Equal(t, 0, len(stdVals))
	assert.Equal(t, 1, len(pulumiVals))
	assert.Equal(t, "\"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3\"", pulumiVals[0])
}

func newTestGenerator(t *testing.T, testFile string) *generator {
	files, err := ioutil.ReadDir(testdataPath)
	if err != nil {
		t.Fatalf("could not read test data: %v", err)
	}

	for _, f := range files {
		if filepath.Base(f.Name()) != testFile {
			continue
		}

		path := filepath.Join(testdataPath, f.Name())
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatalf("could not read %v: %v", path, err)
		}

		parser := syntax.NewParser()
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
		}
		if diags.HasErrors() {
			t.Fatalf("failed to bind program: %v", diags)
		}

		g := &generator{
			program:             program,
			jsonTempSpiller:     &jsonSpiller{},
			ternaryTempSpiller:  &tempSpiller{},
			readDirTempSpiller:  &readDirSpiller{},
			splatSpiller:        &splatSpiller{},
			optionalSpiller:     &optionalSpiller{},
			scopeTraversalRoots: codegen.NewStringSet(),
			arrayHelpers:        make(map[string]*promptToInputArrayHelper),
		}
		g.Formatter = format.NewFormatter(g)
		return g
	}
	t.Fatalf("test file not found")
	return nil
}
