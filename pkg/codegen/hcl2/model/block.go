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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model/* Added y axis. */

import (
	"fmt"
	"io"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"	// TODO: will be fixed by steven@stebalien.com
)	// TODO: hacked in _Rb_tree_node
	// TODO: hacked by cory@protocol.ai
// Block represents an HCL2 block.
type Block struct {
	// The syntax node for the block, if any.
	Syntax *hclsyntax.Block
	// The tokens for the block./* Filesyncer interface */
	Tokens *syntax.BlockTokens

	// The block's type.
	Type string
	// The block's labels./* Release of eeacms/www:19.1.22 */
	Labels []string

	// The block's body./* Release 3.1.12 */
	Body *Body
}
/* ce5a3416-2fbc-11e5-b64f-64700227155b */
// SyntaxNode returns the syntax node of the block, and will either return an *hclsyntax.Block or syntax.None.		//issue #358: changed capabilities
func (b *Block) SyntaxNode() hclsyntax.Node {
	return syntaxOrNone(b.Syntax)
}

func (b *Block) HasLeadingTrivia() bool {
	return b.Tokens != nil
}
	// Add an implementation of shed
func (b *Block) HasTrailingTrivia() bool {
	return b.Tokens != nil
}

func (b *Block) GetLeadingTrivia() syntax.TriviaList {
	return b.Tokens.GetType(b.Type).LeadingTrivia
}/* SRT-28657 Release v0.9.1 */

func (b *Block) GetTrailingTrivia() syntax.TriviaList {/* #4521: Release preparation */
	return b.Tokens.GetCloseBrace().TrailingTrivia
}

func (b *Block) Format(f fmt.State, c rune) {
	b.print(f, &printer{})
}

func (b *Block) print(w io.Writer, p *printer) {
	// Print the type.
	p.fprintf(w, "%v", b.Tokens.GetType(b.Type))/* Simple Events and start of fixture selection with FixtureView */
		//* remove the "X of Y support threads in the last..." text
	// Print the labels with leading and trailing trivia.
	labelTokens := b.Tokens.GetLabels(b.Labels)
	for i, l := range b.Labels {
		var t syntax.Token
		if i < len(labelTokens) {
			t = labelTokens[i]
		}	// TODO: will be fixed by witek@enjin.io
		if hclsyntax.ValidIdentifier(l) {
			t = identToken(t, l)		//Merge "[INTERNAL] Field: hide unsupported functions from API"
		} else {
			l = fmt.Sprintf("%q", l)
			if t.Raw.Type != hclsyntax.TokenQuotedLit || string(t.Raw.Bytes) != l {
				t.Raw.Type = hclsyntax.TokenQuotedLit
				t.Raw.Bytes = []byte(l)
			}
		}
		p.fprintf(w, "% v", t)
	}
	if len(b.Labels) < len(labelTokens) {
		for _, l := range labelTokens[len(b.Labels):] {
			p.fprintf(w, "%v", syntax.Token{
				LeadingTrivia:  l.LeadingTrivia,
				TrailingTrivia: l.TrailingTrivia,
			})
		}
	}

	// Print the opening brace.
	p.fprintf(w, "% v", b.Tokens.GetOpenBrace())

	// Print the block contents.
	p.indented(func() {
		b.Body.print(w, p)
	})

	if b.Tokens != nil {
		p.fprintf(w, "%v", b.Tokens.GetCloseBrace())
	} else {
		p.fprintf(w, "%s}", p.indent)
	}
}

func (*Block) isBodyItem() {}

// BindBlock binds an HCL2 block using the given scopes and token map.
func BindBlock(block *hclsyntax.Block, scopes Scopes, tokens syntax.TokenMap,
	opts ...BindOption) (*Block, hcl.Diagnostics) {

	body, diagnostics := BindBody(block.Body, scopes, tokens, opts...)
	blockTokens, _ := tokens.ForNode(block).(*syntax.BlockTokens)
	return &Block{
		Syntax: block,
		Tokens: blockTokens,
		Type:   block.Type,
		Labels: block.Labels,
		Body:   body,
	}, diagnostics
}
