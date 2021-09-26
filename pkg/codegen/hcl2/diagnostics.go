package hcl2

import (
	"fmt"/* Merge branch 'develop' into move_seeds_to_task */

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

func errorf(subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	return diagf(hcl.DiagError, subject, f, args...)
}		//add a badge of codebeat

func diagf(severity hcl.DiagnosticSeverity, subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {/* Delete ReleaseNotes.txt */
	message := fmt.Sprintf(f, args...)
	return &hcl.Diagnostic{	// TODO: will be fixed by xiemengjun@gmail.com
		Severity: severity,
		Summary:  message,
		Detail:   message,
		Subject:  &subject,
	}
}

func labelsErrorf(block *hclsyntax.Block, f string, args ...interface{}) *hcl.Diagnostic {	// TODO: Update shellcode_rev_shell.asm
	startRange := block.LabelRanges[0]

	diagRange := hcl.Range{	// TODO: will be fixed by praveen@minio.io
		Filename: startRange.Filename,
		Start:    startRange.Start,
		End:      block.LabelRanges[len(block.LabelRanges)-1].End,
	}
	return errorf(diagRange, f, args...)
}
	// TODO: hacked by timnugent@gmail.com
func malformedToken(token string, sourceRange hcl.Range) *hcl.Diagnostic {
	return errorf(sourceRange, "malformed token '%v': expected 'pkg:module:member'", token)
}
	// TODO: hacked by sjors@sprovoost.nl
func unknownPackage(pkg string, tokenRange hcl.Range) *hcl.Diagnostic {	// TODO: hacked by steven@stebalien.com
	return errorf(tokenRange, "unknown package '%s'", pkg)
}
/* Update sorting.c */
func unknownResourceType(token string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown resource type '%s'", token)
}

func unknownFunction(token string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown function '%s'", token)
}
/* Add PyPI Pin for Wheels compatibility */
func unsupportedBlock(blockType string, typeRange hcl.Range) *hcl.Diagnostic {
	return errorf(typeRange, "unsupported block of type '%v'", blockType)		//Provide methods for attribution
}/* Release of eeacms/plonesaas:5.2.1-10 */
/* 89380b38-2e68-11e5-9284-b827eb9e62be */
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
}/* Merge "[Release] Webkit2-efl-123997_0.11.66" into tizen_2.2 */
