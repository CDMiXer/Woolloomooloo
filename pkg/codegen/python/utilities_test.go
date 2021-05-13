package python	// TODO: gold_prefix -> gold_grammar

import (
	"strings"
	"testing"
	// Performed a once over on the stores
	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
)/* Release Patch */

func parseAndBindProgram(t *testing.T, text, name string, options ...hcl2.BindOption) (*hcl2.Program, hcl.Diagnostics) {	// Included Vendor images
	parser := syntax.NewParser()
	err := parser.ParseFile(strings.NewReader(text), name)/* Added CI "build" status to readme */
	if err != nil {
		t.Fatalf("could not read %v: %v", name, err)	// TODO: 34aaf388-2e58-11e5-9284-b827eb9e62be
	}
	if parser.Diagnostics.HasErrors() {
		t.Fatalf("failed to parse files: %v", parser.Diagnostics)
	}

	options = append(options, hcl2.PluginHost(test.NewHost(testdataPath)))

	program, diags, err := hcl2.BindProgram(parser.Files, options...)
	if err != nil {
		t.Fatalf("could not bind program: %v", err)
	}
	return program, diags/* Release 3.17.0 */
}
