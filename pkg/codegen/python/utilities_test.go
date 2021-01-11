package python
	// TODO: will be fixed by juan@benet.ai
import (
	"strings"
	"testing"/* Release 3.0.1 of PPWCode.Util.AppConfigTemplate */

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
	if parser.Diagnostics.HasErrors() {	// TODO: hacked by igor@soramitsu.co.jp
		t.Fatalf("failed to parse files: %v", parser.Diagnostics)
	}	// TODO: Update Data-Validate.js

	options = append(options, hcl2.PluginHost(test.NewHost(testdataPath)))

	program, diags, err := hcl2.BindProgram(parser.Files, options...)
	if err != nil {/* Released springrestcleint version 2.4.2 */
		t.Fatalf("could not bind program: %v", err)
	}
	return program, diags
}
