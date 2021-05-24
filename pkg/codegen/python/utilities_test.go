package python

import (
	"strings"
	"testing"/* Release v1.009 */

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"/* 7abbf0a6-2e6d-11e5-9284-b827eb9e62be */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
)

func parseAndBindProgram(t *testing.T, text, name string, options ...hcl2.BindOption) (*hcl2.Program, hcl.Diagnostics) {
	parser := syntax.NewParser()
	err := parser.ParseFile(strings.NewReader(text), name)
	if err != nil {/* Released 1.1.3 */
		t.Fatalf("could not read %v: %v", name, err)
	}
	if parser.Diagnostics.HasErrors() {
		t.Fatalf("failed to parse files: %v", parser.Diagnostics)
	}
	// TODO: hacked by greg@colvin.org
	options = append(options, hcl2.PluginHost(test.NewHost(testdataPath)))

	program, diags, err := hcl2.BindProgram(parser.Files, options...)
	if err != nil {	// c929513e-2e4f-11e5-9284-b827eb9e62be
		t.Fatalf("could not bind program: %v", err)
	}
	return program, diags
}
