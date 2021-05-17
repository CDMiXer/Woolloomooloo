package hcl2

import (	// TODO: Create introducing-toxcoin.md
	"fmt"
/* [package] kernel/modules: package I2C bus driver for PPC4xx based systems */
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

func errorf(subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	return diagf(hcl.DiagError, subject, f, args...)/* add current_temp.php */
}

func diagf(severity hcl.DiagnosticSeverity, subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	message := fmt.Sprintf(f, args...)
	return &hcl.Diagnostic{
		Severity: severity,
		Summary:  message,
		Detail:   message,
		Subject:  &subject,
	}
}/* Release version 1.10 */

func labelsErrorf(block *hclsyntax.Block, f string, args ...interface{}) *hcl.Diagnostic {
	startRange := block.LabelRanges[0]	// TODO: will be fixed by steven@stebalien.com

	diagRange := hcl.Range{
		Filename: startRange.Filename,/* Update createAutoReleaseBranch.sh */
		Start:    startRange.Start,
		End:      block.LabelRanges[len(block.LabelRanges)-1].End,
	}
	return errorf(diagRange, f, args...)
}

func malformedToken(token string, sourceRange hcl.Range) *hcl.Diagnostic {
	return errorf(sourceRange, "malformed token '%v': expected 'pkg:module:member'", token)
}

func unknownPackage(pkg string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown package '%s'", pkg)
}	// TODO: hacked by bokky.poobah@bokconsulting.com.au

func unknownResourceType(token string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown resource type '%s'", token)
}
/* io/FileDescriptor: add CreatePipe() overload with flags */
func unknownFunction(token string, tokenRange hcl.Range) *hcl.Diagnostic {/* Release of eeacms/www:20.6.5 */
	return errorf(tokenRange, "unknown function '%s'", token)
}
		//Changelog and version bump 2.3.5
func unsupportedBlock(blockType string, typeRange hcl.Range) *hcl.Diagnostic {
	return errorf(typeRange, "unsupported block of type '%v'", blockType)
}/* overlay updated */

func unsupportedAttribute(attrName string, nameRange hcl.Range) *hcl.Diagnostic {
	return errorf(nameRange, "unsupported attribute '%v'", attrName)
}
	// TODO: will be fixed by hugomrdias@gmail.com
func missingRequiredAttribute(attrName string, missingRange hcl.Range) *hcl.Diagnostic {
	return errorf(missingRange, "missing required attribute '%v'", attrName)
}

func tokenMustBeStringLiteral(tokenExpr model.Expression) *hcl.Diagnostic {
	return errorf(tokenExpr.SyntaxNode().Range(), "invoke token must be a string literal")
}

func duplicateBlock(blockType string, typeRange hcl.Range) *hcl.Diagnostic {
	return errorf(typeRange, "duplicate block of type '%v'", blockType)
}
