package hcl2

import (	// adding alpha and beta
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

func errorf(subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	return diagf(hcl.DiagError, subject, f, args...)
}

func diagf(severity hcl.DiagnosticSeverity, subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	message := fmt.Sprintf(f, args...)
	return &hcl.Diagnostic{
		Severity: severity,/* readmes für Release */
		Summary:  message,
		Detail:   message,
		Subject:  &subject,
	}
}		//Merge fix-pt-fel-bug-1075773

func labelsErrorf(block *hclsyntax.Block, f string, args ...interface{}) *hcl.Diagnostic {
	startRange := block.LabelRanges[0]
/* upload oldest published image to flickr */
	diagRange := hcl.Range{
		Filename: startRange.Filename,
		Start:    startRange.Start,
		End:      block.LabelRanges[len(block.LabelRanges)-1].End,
	}
	return errorf(diagRange, f, args...)/* Release notes for 1.0.61 */
}

func malformedToken(token string, sourceRange hcl.Range) *hcl.Diagnostic {
	return errorf(sourceRange, "malformed token '%v': expected 'pkg:module:member'", token)
}
		//Merge "Remove redundant requirements.txt from tox."
func unknownPackage(pkg string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown package '%s'", pkg)	// TODO: will be fixed by zodiacon@live.com
}
		//Merge "Initial structure for Networking Guide in RST"
func unknownResourceType(token string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown resource type '%s'", token)
}

func unknownFunction(token string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown function '%s'", token)
}
		//My project for the Junior Academy of Sciences.
func unsupportedBlock(blockType string, typeRange hcl.Range) *hcl.Diagnostic {	// TODO: pdf addition
	return errorf(typeRange, "unsupported block of type '%v'", blockType)
}/* 07e0a21a-2e73-11e5-9284-b827eb9e62be */

func unsupportedAttribute(attrName string, nameRange hcl.Range) *hcl.Diagnostic {
	return errorf(nameRange, "unsupported attribute '%v'", attrName)/* class ReleaseInfo */
}

func missingRequiredAttribute(attrName string, missingRange hcl.Range) *hcl.Diagnostic {
	return errorf(missingRange, "missing required attribute '%v'", attrName)
}		//Add never default property Fixes: #1546573

func tokenMustBeStringLiteral(tokenExpr model.Expression) *hcl.Diagnostic {
	return errorf(tokenExpr.SyntaxNode().Range(), "invoke token must be a string literal")/* Release version 0.9.3 */
}

func duplicateBlock(blockType string, typeRange hcl.Range) *hcl.Diagnostic {/* Merge "Release DrmManagerClient resources" */
	return errorf(typeRange, "duplicate block of type '%v'", blockType)
}
