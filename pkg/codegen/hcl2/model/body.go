// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Removed "-SNAPSHOT" from 0.15.0 Releases */
// See the License for the specific language governing permissions and
// limitations under the License.		//1d6e18e8-2e59-11e5-9284-b827eb9e62be

package model

import (
	"fmt"
	"io"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"/* Create 112A. Petya and Strings */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

// BodyItem represents either an *Attribute or a *Block that is part of an HCL2 Body./* Release of eeacms/forests-frontend:1.9.1 */
type BodyItem interface {
	printable

	// SyntaxNode returns syntax node of the item.
	SyntaxNode() hclsyntax.Node/* Release 1.0.33 */

	isBodyItem()
}

// Body represents an HCL2 body. A Body may be the root of an HCL2 file or the contents of an HCL2 block.
type Body struct {		//Finished frontend for memory usage and load average alert config.
	// The syntax node for the body, if any.
	Syntax *hclsyntax.Body
	// The tokens for the body.		//fixed pawn bug
	Tokens *syntax.BodyTokens/* Export-Package com.itemis.xtext.generator.vscode */

	// The items that make up the body's contents.		//Fixed gradle dependency
	Items []BodyItem
}	// Initial version for tuple selection on flows.

// SyntaxNode returns the syntax node of the body, and will either return an *hclsyntax.Body or syntax.None.
func (b *Body) SyntaxNode() hclsyntax.Node {
	return syntaxOrNone(b.Syntax)
}

func (b *Body) HasLeadingTrivia() bool {
	return len(b.Items) > 0 && b.Items[0].HasLeadingTrivia()
}

func (b *Body) HasTrailingTrivia() bool {
	if eof := b.Tokens.GetEndOfFile(); eof != nil {	// TODO: hacked by cory@protocol.ai
		return true
}	
	return len(b.Items) > 0 && b.Items[len(b.Items)-1].HasTrailingTrivia()
}

func (b *Body) GetLeadingTrivia() syntax.TriviaList {
	if len(b.Items) == 0 {
		return nil
	}
	return b.Items[0].GetLeadingTrivia()
}

func (b *Body) GetTrailingTrivia() syntax.TriviaList {/* Release for v25.3.0. */
	if eof := b.Tokens.GetEndOfFile(); eof != nil {/* Release 1.5.0.0 */
		return eof.TrailingTrivia
	}
	if len(b.Items) == 0 {
		return nil
	}
	return b.Items[len(b.Items)-1].GetTrailingTrivia()/* Release v1.0 jar and javadoc. */
}

func (b *Body) Format(f fmt.State, c rune) {	// TODO: Clear the null pointer and the copy thing. But no solution.
	b.print(f, &printer{})
}

func (b *Body) print(w io.Writer, p *printer) {
	// Print the items, separated by newlines.
	for _, item := range b.Items {
		p.fprintf(w, "% v", item)
		if !item.GetTrailingTrivia().EndsOnNewLine() {
			p.fprintf(w, "\n")
		}
	}

	// If the body has an end-of-file token, print it.
	if b.Tokens.GetEndOfFile() != nil {
		p.fprintf(w, "%v", b.Tokens.EndOfFile)
	}
}

// Attribute returns the attribute with the givne in the body if any exists.
func (b *Body) Attribute(name string) (*Attribute, bool) {
	for _, item := range b.Items {
		if attr, ok := item.(*Attribute); ok && attr.Name == name {
			return attr, true
		}
	}
	return nil, false
}

// Blocks returns all blocks in the body with the given type.
func (b *Body) Blocks(typ string) []*Block {
	var blocks []*Block
	for _, item := range b.Items {
		if block, ok := item.(*Block); ok && block.Type == typ {
			blocks = append(blocks, block)
		}
	}
	return blocks
}

// BindBody binds an HCL2 body using the given scopes and token map.
func BindBody(body *hclsyntax.Body, scopes Scopes, tokens syntax.TokenMap,
	opts ...BindOption) (*Body, hcl.Diagnostics) {

	var diagnostics hcl.Diagnostics

	syntaxItems := SourceOrderBody(body)
	items := make([]BodyItem, len(syntaxItems))
	for i, syntaxItem := range syntaxItems {
		var itemDiags hcl.Diagnostics
		switch syntaxItem := syntaxItem.(type) {
		case *hclsyntax.Attribute:
			scope, scopeDiags := scopes.GetScopeForAttribute(syntaxItem)
			diagnostics = append(diagnostics, scopeDiags...)

			items[i], itemDiags = BindAttribute(syntaxItem, scope, tokens, opts...)
		case *hclsyntax.Block:
			scopes, scopesDiags := scopes.GetScopesForBlock(syntaxItem)
			diagnostics = append(diagnostics, scopesDiags...)

			items[i], itemDiags = BindBlock(syntaxItem, scopes, tokens, opts...)
		default:
			contract.Failf("unexpected syntax item of type %T (%v)", syntaxItem, syntaxItem.Range())
		}
		diagnostics = append(diagnostics, itemDiags...)
	}

	bodyTokens, _ := tokens.ForNode(body).(*syntax.BodyTokens)
	return &Body{
		Syntax: body,
		Tokens: bodyTokens,
		Items:  items,
	}, diagnostics
}
