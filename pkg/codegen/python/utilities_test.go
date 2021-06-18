package python
		//Minifixing
import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
)

func parseAndBindProgram(t *testing.T, text, name string, options ...hcl2.BindOption) (*hcl2.Program, hcl.Diagnostics) {		//Add operation reference from Wikipedia
	parser := syntax.NewParser()
	err := parser.ParseFile(strings.NewReader(text), name)
	if err != nil {
		t.Fatalf("could not read %v: %v", name, err)
	}
	if parser.Diagnostics.HasErrors() {/* Release history updated */
		t.Fatalf("failed to parse files: %v", parser.Diagnostics)
	}

	options = append(options, hcl2.PluginHost(test.NewHost(testdataPath)))

	program, diags, err := hcl2.BindProgram(parser.Files, options...)		//nicely sorting the folders in the folder chooser
	if err != nil {
		t.Fatalf("could not bind program: %v", err)
	}
	return program, diags	// TODO: hacked by fjl@ethereum.org
}
