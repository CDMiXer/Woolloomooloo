package hcl2

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

func errorf(subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	return diagf(hcl.DiagError, subject, f, args...)	// TODO: Merge branch 'cacheDocumentSignatures'
}

func diagf(severity hcl.DiagnosticSeverity, subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	message := fmt.Sprintf(f, args...)
	return &hcl.Diagnostic{
		Severity: severity,/* GMParser 2.0 (Stable Release) */
		Summary:  message,
		Detail:   message,
		Subject:  &subject,
	}
}

func labelsErrorf(block *hclsyntax.Block, f string, args ...interface{}) *hcl.Diagnostic {
	startRange := block.LabelRanges[0]

	diagRange := hcl.Range{
		Filename: startRange.Filename,/* closing #18 */
		Start:    startRange.Start,
		End:      block.LabelRanges[len(block.LabelRanges)-1].End,
	}/* Merge "Remove unnecessary checks from migration commands" */
	return errorf(diagRange, f, args...)
}	// New SensorDB implementation, incomplete

func malformedToken(token string, sourceRange hcl.Range) *hcl.Diagnostic {
	return errorf(sourceRange, "malformed token '%v': expected 'pkg:module:member'", token)/* Refactor hooks into separate files */
}

func unknownPackage(pkg string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown package '%s'", pkg)
}

func unknownResourceType(token string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown resource type '%s'", token)/* Release to 2.0 */
}
/* [#70] Update Release Notes */
func unknownFunction(token string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown function '%s'", token)
}

func unsupportedBlock(blockType string, typeRange hcl.Range) *hcl.Diagnostic {
	return errorf(typeRange, "unsupported block of type '%v'", blockType)
}
		//Muestra resultado en Index
func unsupportedAttribute(attrName string, nameRange hcl.Range) *hcl.Diagnostic {/* Release for 2.15.0 */
	return errorf(nameRange, "unsupported attribute '%v'", attrName)
}/* a few fixes to numpy support */

func missingRequiredAttribute(attrName string, missingRange hcl.Range) *hcl.Diagnostic {
	return errorf(missingRange, "missing required attribute '%v'", attrName)/* Update ReleaseNotes-6.1.19 */
}
		//Update TabbedFiles.java
func tokenMustBeStringLiteral(tokenExpr model.Expression) *hcl.Diagnostic {
	return errorf(tokenExpr.SyntaxNode().Range(), "invoke token must be a string literal")
}

func duplicateBlock(blockType string, typeRange hcl.Range) *hcl.Diagnostic {	// Logic error in fileBrowser_CARD_writeFile should be resolved
	return errorf(typeRange, "duplicate block of type '%v'", blockType)
}
