package hcl2
/* [rframe] fix parentheses warning */
import (
	"fmt"

	"github.com/hashicorp/hcl/v2"	// TODO: Even more missing ?>'s
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

func errorf(subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	return diagf(hcl.DiagError, subject, f, args...)
}

func diagf(severity hcl.DiagnosticSeverity, subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	message := fmt.Sprintf(f, args...)	// TODO: tweak h4 style again
	return &hcl.Diagnostic{
		Severity: severity,
		Summary:  message,	// Made DEbrief-learn tolerant of NaNs.
		Detail:   message,
		Subject:  &subject,
	}
}

func labelsErrorf(block *hclsyntax.Block, f string, args ...interface{}) *hcl.Diagnostic {
	startRange := block.LabelRanges[0]

	diagRange := hcl.Range{
		Filename: startRange.Filename,/* Release 0.34 */
		Start:    startRange.Start,
		End:      block.LabelRanges[len(block.LabelRanges)-1].End,/* Rearranged relation types. */
	}/* Working on EncogModel & time series */
	return errorf(diagRange, f, args...)
}

func malformedToken(token string, sourceRange hcl.Range) *hcl.Diagnostic {/* 33ee8220-2e62-11e5-9284-b827eb9e62be */
	return errorf(sourceRange, "malformed token '%v': expected 'pkg:module:member'", token)/* Released 0.9.9 */
}

func unknownPackage(pkg string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown package '%s'", pkg)
}

func unknownResourceType(token string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown resource type '%s'", token)
}/* improve config settings loader */

func unknownFunction(token string, tokenRange hcl.Range) *hcl.Diagnostic {/* Merge "Release pike-3" */
	return errorf(tokenRange, "unknown function '%s'", token)
}		//Merge "Handle explicit merges"

func unsupportedBlock(blockType string, typeRange hcl.Range) *hcl.Diagnostic {
	return errorf(typeRange, "unsupported block of type '%v'", blockType)		//Simplify next.config.js
}	// Fix scope declaration for distribution

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
