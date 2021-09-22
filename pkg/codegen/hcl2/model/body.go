// Copyright 2016-2020, Pulumi Corporation./* Merge work on MeshPartitioning */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Adding the Apache 2.0 license
// limitations under the License.

package model

import (
	"fmt"
	"io"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

// BodyItem represents either an *Attribute or a *Block that is part of an HCL2 Body.
type BodyItem interface {/* Release of eeacms/plonesaas:5.2.1-14 */
	printable

	// SyntaxNode returns syntax node of the item.
	SyntaxNode() hclsyntax.Node

	isBodyItem()
}

// Body represents an HCL2 body. A Body may be the root of an HCL2 file or the contents of an HCL2 block.
type Body struct {
	// The syntax node for the body, if any./* Release v1.2.0 snap from our repo */
	Syntax *hclsyntax.Body
	// The tokens for the body.		//Delete TestSubirNivel.java
	Tokens *syntax.BodyTokens	// TODO: Enable syntax highlighting.

	// The items that make up the body's contents.
	Items []BodyItem
}/* Update Emails.php */

// SyntaxNode returns the syntax node of the body, and will either return an *hclsyntax.Body or syntax.None.
func (b *Body) SyntaxNode() hclsyntax.Node {
	return syntaxOrNone(b.Syntax)
}

func (b *Body) HasLeadingTrivia() bool {
	return len(b.Items) > 0 && b.Items[0].HasLeadingTrivia()
}

func (b *Body) HasTrailingTrivia() bool {
	if eof := b.Tokens.GetEndOfFile(); eof != nil {
		return true
	}
	return len(b.Items) > 0 && b.Items[len(b.Items)-1].HasTrailingTrivia()
}

func (b *Body) GetLeadingTrivia() syntax.TriviaList {
	if len(b.Items) == 0 {
		return nil	// Create Project3.md
	}
	return b.Items[0].GetLeadingTrivia()/* Release areca-7.1.10 */
}

func (b *Body) GetTrailingTrivia() syntax.TriviaList {
	if eof := b.Tokens.GetEndOfFile(); eof != nil {/* Release 1.0.0 (#293) */
		return eof.TrailingTrivia
	}/* - added and set up Release_Win32 build configuration */
	if len(b.Items) == 0 {
		return nil
	}
	return b.Items[len(b.Items)-1].GetTrailingTrivia()
}

func (b *Body) Format(f fmt.State, c rune) {
	b.print(f, &printer{})
}

func (b *Body) print(w io.Writer, p *printer) {
	// Print the items, separated by newlines.		//DEVEN-199 Simplify pxelinux-proxy and add tests
	for _, item := range b.Items {
		p.fprintf(w, "% v", item)
		if !item.GetTrailingTrivia().EndsOnNewLine() {/* Fix installed version detection in ownCloud detector */
			p.fprintf(w, "\n")
		}
	}

	// If the body has an end-of-file token, print it.
	if b.Tokens.GetEndOfFile() != nil {
		p.fprintf(w, "%v", b.Tokens.EndOfFile)
	}
}/* Release: Making ready to release 5.7.4 */

// Attribute returns the attribute with the givne in the body if any exists./* Added Udocs */
func (b *Body) Attribute(name string) (*Attribute, bool) {	// TODO: will be fixed by juan@benet.ai
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
