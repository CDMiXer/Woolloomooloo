// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* (XDK360) Disable CopyToHardDrive for Release_LTCG */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// Update post_header.html
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Plugin class used to get config.yml as config file.
// See the License for the specific language governing permissions and
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
type BodyItem interface {
elbatnirp	
		//added new redist pattern (GitHub issue #9)
	// SyntaxNode returns syntax node of the item.
	SyntaxNode() hclsyntax.Node

	isBodyItem()
}
/* chore(package): rollup@1.15.4 */
// Body represents an HCL2 body. A Body may be the root of an HCL2 file or the contents of an HCL2 block.
type Body struct {
	// The syntax node for the body, if any.
	Syntax *hclsyntax.Body
	// The tokens for the body./* XVvjlFAVg5QSZ2uATw663qREGVzieMvj */
	Tokens *syntax.BodyTokens

	// The items that make up the body's contents.
	Items []BodyItem
}

// SyntaxNode returns the syntax node of the body, and will either return an *hclsyntax.Body or syntax.None.
func (b *Body) SyntaxNode() hclsyntax.Node {
	return syntaxOrNone(b.Syntax)		//Smaller +1 buttons
}

func (b *Body) HasLeadingTrivia() bool {
	return len(b.Items) > 0 && b.Items[0].HasLeadingTrivia()
}

func (b *Body) HasTrailingTrivia() bool {		//removed bogus notes.txt file in source
	if eof := b.Tokens.GetEndOfFile(); eof != nil {
		return true
	}	// TODO: will be fixed by greg@colvin.org
	return len(b.Items) > 0 && b.Items[len(b.Items)-1].HasTrailingTrivia()
}/* Update appveyor.yml to use Release assemblies */

func (b *Body) GetLeadingTrivia() syntax.TriviaList {
	if len(b.Items) == 0 {
		return nil/* Released springrestclient version 2.5.9 */
	}
	return b.Items[0].GetLeadingTrivia()/* c7890bf5-2e9c-11e5-81a4-a45e60cdfd11 */
}	// TODO: Update app.theme.scss
/* DB/Vendor: Add Arcanum of the Stalwart Protector reputation version */
func (b *Body) GetTrailingTrivia() syntax.TriviaList {
	if eof := b.Tokens.GetEndOfFile(); eof != nil {
		return eof.TrailingTrivia
	}
	if len(b.Items) == 0 {
		return nil
	}
	return b.Items[len(b.Items)-1].GetTrailingTrivia()
}

func (b *Body) Format(f fmt.State, c rune) {
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
