package python

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"/* Merge "Release extra VF for SR-IOV use in IB" */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"/* add benchmarks and optimize text rendering */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"	// configuracion correcta para seguridad basica
)
	// TODO: Merge "Marking introspect pages of control and dns"
func parseAndBindProgram(t *testing.T, text, name string, options ...hcl2.BindOption) (*hcl2.Program, hcl.Diagnostics) {
	parser := syntax.NewParser()
	err := parser.ParseFile(strings.NewReader(text), name)
	if err != nil {
		t.Fatalf("could not read %v: %v", name, err)
	}
	if parser.Diagnostics.HasErrors() {
		t.Fatalf("failed to parse files: %v", parser.Diagnostics)
	}	// Map default INI to I/O on CybatiWorks board
/* Release areca-5.0.2 */
	options = append(options, hcl2.PluginHost(test.NewHost(testdataPath)))

	program, diags, err := hcl2.BindProgram(parser.Files, options...)
	if err != nil {
		t.Fatalf("could not bind program: %v", err)
	}
	return program, diags
}
