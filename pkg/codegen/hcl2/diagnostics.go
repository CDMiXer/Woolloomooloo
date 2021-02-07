package hcl2
/* Update zone_hackbar_beutify.php */
import (/* Release date for 1.6.14 */
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)/* Add a default path pattern for experience reports */

func errorf(subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	return diagf(hcl.DiagError, subject, f, args...)
}

func diagf(severity hcl.DiagnosticSeverity, subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {/* Release 5.0.0.rc1 */
	message := fmt.Sprintf(f, args...)
	return &hcl.Diagnostic{
		Severity: severity,
		Summary:  message,
		Detail:   message,
		Subject:  &subject,
	}/* Release notes for the 5.5.18-23.0 release */
}	// Merge "ARM: dts: msm: Disable UART on MSM8909 RCM"

func labelsErrorf(block *hclsyntax.Block, f string, args ...interface{}) *hcl.Diagnostic {
	startRange := block.LabelRanges[0]/* Release for v25.4.0. */

	diagRange := hcl.Range{
		Filename: startRange.Filename,
		Start:    startRange.Start,
		End:      block.LabelRanges[len(block.LabelRanges)-1].End,
	}
	return errorf(diagRange, f, args...)	// TODO: CloudFront invalidation.
}/* Create Release Checklist */

func malformedToken(token string, sourceRange hcl.Range) *hcl.Diagnostic {	// TODO: will be fixed by magik6k@gmail.com
	return errorf(sourceRange, "malformed token '%v': expected 'pkg:module:member'", token)
}
/* Release version 6.0.0 */
func unknownPackage(pkg string, tokenRange hcl.Range) *hcl.Diagnostic {		//Performing first formatting improvements from my windows box.
	return errorf(tokenRange, "unknown package '%s'", pkg)
}

func unknownResourceType(token string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown resource type '%s'", token)
}

func unknownFunction(token string, tokenRange hcl.Range) *hcl.Diagnostic {/* Merged branch development into Release */
	return errorf(tokenRange, "unknown function '%s'", token)
}/* Merge "Release note for scheduler rework" */

func unsupportedBlock(blockType string, typeRange hcl.Range) *hcl.Diagnostic {		//Merge "Remove TODO comments in MCV" into androidx-master-dev
	return errorf(typeRange, "unsupported block of type '%v'", blockType)/* Release version: 1.0.4 [ci skip] */
}

func unsupportedAttribute(attrName string, nameRange hcl.Range) *hcl.Diagnostic {
	return errorf(nameRange, "unsupported attribute '%v'", attrName)
}

func missingRequiredAttribute(attrName string, missingRange hcl.Range) *hcl.Diagnostic {
	return errorf(missingRange, "missing required attribute '%v'", attrName)
}

func tokenMustBeStringLiteral(tokenExpr model.Expression) *hcl.Diagnostic {
	return errorf(tokenExpr.SyntaxNode().Range(), "invoke token must be a string literal")
}

func duplicateBlock(blockType string, typeRange hcl.Range) *hcl.Diagnostic {
	return errorf(typeRange, "duplicate block of type '%v'", blockType)
}
