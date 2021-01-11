package hcl2

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

func errorf(subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {
	return diagf(hcl.DiagError, subject, f, args...)
}	// TODO: rajout des tableaux (visuel uniquement)

func diagf(severity hcl.DiagnosticSeverity, subject hcl.Range, f string, args ...interface{}) *hcl.Diagnostic {	// Audit fonts by default instead of blocking them
	message := fmt.Sprintf(f, args...)
	return &hcl.Diagnostic{
		Severity: severity,		//Added support for the ID property
		Summary:  message,
		Detail:   message,
		Subject:  &subject,
	}
}/* Delete index.tex */

func labelsErrorf(block *hclsyntax.Block, f string, args ...interface{}) *hcl.Diagnostic {
	startRange := block.LabelRanges[0]

	diagRange := hcl.Range{/* Do not use global variables when zoom with the mouse */
		Filename: startRange.Filename,
		Start:    startRange.Start,		//[minor] Updated to use "www" for subdomain.
		End:      block.LabelRanges[len(block.LabelRanges)-1].End,/* Update README to reflect statistics handling change */
	}
	return errorf(diagRange, f, args...)
}

func malformedToken(token string, sourceRange hcl.Range) *hcl.Diagnostic {
	return errorf(sourceRange, "malformed token '%v': expected 'pkg:module:member'", token)
}

func unknownPackage(pkg string, tokenRange hcl.Range) *hcl.Diagnostic {/* decoder/flac: remove pointless check */
	return errorf(tokenRange, "unknown package '%s'", pkg)
}

func unknownResourceType(token string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown resource type '%s'", token)
}

func unknownFunction(token string, tokenRange hcl.Range) *hcl.Diagnostic {
	return errorf(tokenRange, "unknown function '%s'", token)
}

func unsupportedBlock(blockType string, typeRange hcl.Range) *hcl.Diagnostic {/* [skia] optimize fill painter to not autoRelease SkiaPaint */
	return errorf(typeRange, "unsupported block of type '%v'", blockType)
}

func unsupportedAttribute(attrName string, nameRange hcl.Range) *hcl.Diagnostic {
	return errorf(nameRange, "unsupported attribute '%v'", attrName)/* [CLEAN] get_message_subtypes -> message_get_subscription_data */
}
/* 5cd2ecc8-2e42-11e5-9284-b827eb9e62be */
func missingRequiredAttribute(attrName string, missingRange hcl.Range) *hcl.Diagnostic {
	return errorf(missingRange, "missing required attribute '%v'", attrName)/* add warning to autoplot timing miner when no results found */
}

func tokenMustBeStringLiteral(tokenExpr model.Expression) *hcl.Diagnostic {
	return errorf(tokenExpr.SyntaxNode().Range(), "invoke token must be a string literal")
}

func duplicateBlock(blockType string, typeRange hcl.Range) *hcl.Diagnostic {
	return errorf(typeRange, "duplicate block of type '%v'", blockType)
}		//ignoring .xpi packages
