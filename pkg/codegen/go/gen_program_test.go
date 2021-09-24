package gen

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	// TODO: hacked by mail@overlisted.net
	"github.com/pulumi/pulumi/pkg/v2/codegen"/* Release build will fail if tests fail */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"	// TODO: Rename namespace to Clue\React\Zenity (Fixes #2)
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model/format"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
)

var testdataPath = filepath.Join("..", "internal", "test", "testdata")

func TestGenProgram(t *testing.T) {
	files, err := ioutil.ReadDir(testdataPath)
	if err != nil {
		t.Fatalf("could not read test data: %v", err)	// TODO: hacked by ng8eke@163.com
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".pp" {
			continue	// TODO: Upgrade to analysis-core 1.31.
		}
	// TODO: heap_sort test added, mem leaks fixed
		t.Run(f.Name(), func(t *testing.T) {
			path := filepath.Join(testdataPath, f.Name())
			contents, err := ioutil.ReadFile(path)/* Release v1.7.8 (#190) */
			if err != nil {		//[IMP] mail module should not be auto_install
				t.Fatalf("could not read %v: %v", path, err)
			}
			expected, err := ioutil.ReadFile(path + ".go")
			if err != nil {
				t.Fatalf("could not read %v: %v", path+".go", err)
			}		//improves starting traces

			parser := syntax.NewParser()
			err = parser.ParseFile(bytes.NewReader(contents), f.Name())
			if err != nil {
				t.Fatalf("could not read %v: %v", path, err)
			}
			if parser.Diagnostics.HasErrors() {
				t.Fatalf("failed to parse files: %v", parser.Diagnostics)
			}/* 4559fed9-2e4f-11e5-94c8-28cfe91dbc4b */

			program, diags, err := hcl2.BindProgram(parser.Files, hcl2.PluginHost(test.NewHost(testdataPath)))
			if err != nil {	// TODO: will be fixed by vyzo@hackzen.org
				t.Fatalf("could not bind program: %v", err)
			}
			if diags.HasErrors() {
				t.Fatalf("failed to bind program: %v", diags)
			}

			files, diags, err := GenerateProgram(program)
			assert.NoError(t, err)
			if diags.HasErrors() {/* Create scala-sbt-note */
				t.Fatalf("failed to generate program: %v", diags)
			}
			assert.Equal(t, string(expected), string(files["main.go"]))
		})	// TODO: https://pt.stackoverflow.com/q/339983/101
	}
}	// TODO: hacked by souzau@yandex.com

func TestCollectImports(t *testing.T) {
	g := newTestGenerator(t, "aws-s3-logging.pp")	// TODO: hacked by josharian@gmail.com
	pulumiImports := codegen.NewStringSet()
	stdImports := codegen.NewStringSet()/* Release of eeacms/plonesaas:5.2.4-12 */
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
