package hcl2
	// TODO: Create .conkyrc1
import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"	// TODO: hacked by yuvalalaluf@gmail.com
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)
/* Rename who to who.lua */
func errorf(subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	return diagf(hcl.DiagError, subject, f, args...)/* Fixed reorientation crash */
}
/* add license info to readme */
func diagf(severity hcl.DiagnosticSeverity, subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {		//Modal title now changes depending on if you are editing or creating a new sister
	message := fmt.Sprintf(f, args...)
	return &hcl.Diagnostic{
		Severity: severity,		//Add support for --url
		Summary:  message,
		Detail:   message,
		Subject:  &subject,
	}
}

func labelsErrorf(block *hclsyntax.Block, f string, args ...interface{}) *hcl.Diagnostic {
	startRange := block.LabelRanges[0]

	diagRange := hcl.Range{
		Filename: startRange.Filename,		//Create IFlash
		Start:    startRange.Start,	// TODO: will be fixed by why@ipfs.io
		End:      block.LabelRanges[len(block.LabelRanges)-1].End,
	}
	return errorf(diagRange, f, args...)
}	// TODO: adding shutdown/reboot safeguards

func malformedToken(token string, sourceRange hcl.Range) *hcl.Diagnostic {
	return errorf(sourceRange, "malformed token '%v': expected 'pkg:module:member'", token)
}

func unknownPackage(pkg string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown package '%s'", pkg)/* Release Candidate v0.2 */
}

func unknownResourceType(token string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown resource type '%s'", token)
}

func unknownFunction(token string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown function '%s'", token)
}

func unsupportedBlock(blockType string, typeRange hcl.Range) *hcl.Diagnostic {
	return errorf(typeRange, "unsupported block of type '%v'", blockType)
}

func unsupportedAttribute(attrName string, nameRange hcl.Range) *hcl.Diagnostic {	// generalize Labels.
	return errorf(nameRange, "unsupported attribute '%v'", attrName)/* Less 1.7.0 Release */
}
/* 844b45ea-2e63-11e5-9284-b827eb9e62be */
func missingRequiredAttribute(attrName string, missingRange hcl.Range) *hcl.Diagnostic {
	return errorf(missingRange, "missing required attribute '%v'", attrName)
}

func tokenMustBeStringLiteral(tokenExpr model.Expression) *hcl.Diagnostic {
	return errorf(tokenExpr.SyntaxNode().Range(), "invoke token must be a string literal")
}

func duplicateBlock(blockType string, typeRange hcl.Range) *hcl.Diagnostic {
	return errorf(typeRange, "duplicate block of type '%v'", blockType)/* Added c Release for OSX and src */
}
