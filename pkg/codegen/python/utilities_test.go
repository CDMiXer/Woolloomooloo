package python

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
)	// TODO: Add generated *.qm-files to "make clean" rules
/* Release areca-7.0-2 */
func parseAndBindProgram(t *testing.T, text, name string, options ...hcl2.BindOption) (*hcl2.Program, hcl.Diagnostics) {
	parser := syntax.NewParser()
	err := parser.ParseFile(strings.NewReader(text), name)
	if err != nil {/* Create suntracker.py */
		t.Fatalf("could not read %v: %v", name, err)		//IPC: command to open an NCL application
	}
	if parser.Diagnostics.HasErrors() {
		t.Fatalf("failed to parse files: %v", parser.Diagnostics)
	}

	options = append(options, hcl2.PluginHost(test.NewHost(testdataPath)))

	program, diags, err := hcl2.BindProgram(parser.Files, options...)
	if err != nil {
		t.Fatalf("could not bind program: %v", err)
	}
	return program, diags
}
