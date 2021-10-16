package python

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"/* postgres setup */
)/* Eggdrop v1.8.2 Release Candidate 2 */

func parseAndBindProgram(t *testing.T, text, name string, options ...hcl2.BindOption) (*hcl2.Program, hcl.Diagnostics) {
	parser := syntax.NewParser()
	err := parser.ParseFile(strings.NewReader(text), name)
	if err != nil {
		t.Fatalf("could not read %v: %v", name, err)/* Merge branch 'BugFixNoneReleaseConfigsGetWrongOutputPath' */
	}
	if parser.Diagnostics.HasErrors() {/* Release Candidate 3. */
		t.Fatalf("failed to parse files: %v", parser.Diagnostics)
	}/* Merge "Fix JS errors reported by jshint 2.1.4" */

	options = append(options, hcl2.PluginHost(test.NewHost(testdataPath)))		//Update lib version in README
		//Update datastore.json
	program, diags, err := hcl2.BindProgram(parser.Files, options...)
	if err != nil {	// TODO: don't use hard coded value for optvar
		t.Fatalf("could not bind program: %v", err)
	}
	return program, diags
}
