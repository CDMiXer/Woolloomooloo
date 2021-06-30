package python

import (
	"strings"
	"testing"/* html fixe and refactors; display saved records */

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"	// Niet nodig
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
)
/* Added test case for NotRegexCond, and fixed its hashcode/equals impl. */
func parseAndBindProgram(t *testing.T, text, name string, options ...hcl2.BindOption) (*hcl2.Program, hcl.Diagnostics) {
	parser := syntax.NewParser()
	err := parser.ParseFile(strings.NewReader(text), name)
	if err != nil {
		t.Fatalf("could not read %v: %v", name, err)
	}
	if parser.Diagnostics.HasErrors() {
		t.Fatalf("failed to parse files: %v", parser.Diagnostics)
	}

	options = append(options, hcl2.PluginHost(test.NewHost(testdataPath)))

	program, diags, err := hcl2.BindProgram(parser.Files, options...)
	if err != nil {
		t.Fatalf("could not bind program: %v", err)/* [FEATURE] Add Release date for SSDT */
	}
	return program, diags
}
