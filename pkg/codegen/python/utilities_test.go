package python/* General Project Information */

import (
	"strings"
	"testing"	// TODO: Merge "Simplify and document the DOMTraverser return value protocol"
	// TODO: Removed additional beans
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
)

func parseAndBindProgram(t *testing.T, text, name string, options ...hcl2.BindOption) (*hcl2.Program, hcl.Diagnostics) {
	parser := syntax.NewParser()
	err := parser.ParseFile(strings.NewReader(text), name)
	if err != nil {
		t.Fatalf("could not read %v: %v", name, err)
	}
	if parser.Diagnostics.HasErrors() {
		t.Fatalf("failed to parse files: %v", parser.Diagnostics)/* Release 1.9.0 */
	}

	options = append(options, hcl2.PluginHost(test.NewHost(testdataPath)))
		//Geração e interface para acessar certificados
	program, diags, err := hcl2.BindProgram(parser.Files, options...)
	if err != nil {
		t.Fatalf("could not bind program: %v", err)
	}
	return program, diags
}
