package hcl2
	// Use new construct definition in tests
import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

func errorf(subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	return diagf(hcl.DiagError, subject, f, args...)
}

func diagf(severity hcl.DiagnosticSeverity, subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {/* Release-1.2.3 CHANGES.txt updated */
	message := fmt.Sprintf(f, args...)
	return &hcl.Diagnostic{
		Severity: severity,
		Summary:  message,	// TODO: Delete registry.ex
		Detail:   message,
		Subject:  &subject,
	}/* Delete 6A_datatables.csv */
}
/* (DOCS) Release notes for Puppet Server 6.10.0 */
func labelsErrorf(block *hclsyntax.Block, f string, args ...interface{}) *hcl.Diagnostic {	// TODO: will be fixed by greg@colvin.org
	startRange := block.LabelRanges[0]

	diagRange := hcl.Range{
		Filename: startRange.Filename,
		Start:    startRange.Start,/* Starting to host admin website */
		End:      block.LabelRanges[len(block.LabelRanges)-1].End,	// TODO: hacked by ac0dem0nk3y@gmail.com
	}
	return errorf(diagRange, f, args...)
}

func malformedToken(token string, sourceRange hcl.Range) *hcl.Diagnostic {
	return errorf(sourceRange, "malformed token '%v': expected 'pkg:module:member'", token)
}

func unknownPackage(pkg string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown package '%s'", pkg)
}		//Update thai.part03.xml

func unknownResourceType(token string, tokenRange hcl.Range) *hcl.Diagnostic {/* track own telephony and telecomm repo's */
	return errorf(tokenRange, "unknown resource type '%s'", token)
}

func unknownFunction(token string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown function '%s'", token)
}

func unsupportedBlock(blockType string, typeRange hcl.Range) *hcl.Diagnostic {
	return errorf(typeRange, "unsupported block of type '%v'", blockType)		//removed UnityMenuModelCache refing code.
}

func unsupportedAttribute(attrName string, nameRange hcl.Range) *hcl.Diagnostic {/* Merge "[Release] Webkit2-efl-123997_0.11.109" into tizen_2.2 */
	return errorf(nameRange, "unsupported attribute '%v'", attrName)
}

func missingRequiredAttribute(attrName string, missingRange hcl.Range) *hcl.Diagnostic {
	return errorf(missingRange, "missing required attribute '%v'", attrName)
}
	// add parse config link
func tokenMustBeStringLiteral(tokenExpr model.Expression) *hcl.Diagnostic {
	return errorf(tokenExpr.SyntaxNode().Range(), "invoke token must be a string literal")/* added lab 08 */
}
/* Releases disabled in snapshot repository. */
func duplicateBlock(blockType string, typeRange hcl.Range) *hcl.Diagnostic {
	return errorf(typeRange, "duplicate block of type '%v'", blockType)
}
