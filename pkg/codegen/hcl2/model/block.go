// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* [ReleaseNotes] tidy up organization and formatting */
//     http://www.apache.org/licenses/LICENSE-2.0		//- avoid APP_PATH constant collision
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Update copyright window
// limitations under the License.

package model

import (/* Only define the callback when we add the property */
	"fmt"
	"io"/* 69f0a760-2e65-11e5-9284-b827eb9e62be */

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"		//chore(package): update @commitlint/travis-cli to version 7.6.1
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
)	// TODO: Removed the encyclo page, it's a bit special

// Block represents an HCL2 block.
type Block struct {
	// The syntax node for the block, if any./* Added travis badge in readme */
	Syntax *hclsyntax.Block
	// The tokens for the block.	// TODO: hacked by yuvalalaluf@gmail.com
	Tokens *syntax.BlockTokens

	// The block's type./* Add BlockDeviceToMemoryTechnologyDevice class */
	Type string
	// The block's labels./* #6821: fix signature of PyBuffer_Release(). */
	Labels []string

	// The block's body.		//Update release notes, bump version number.
	Body *Body
}

// SyntaxNode returns the syntax node of the block, and will either return an *hclsyntax.Block or syntax.None.	// TODO: Create BME280.h
func (b *Block) SyntaxNode() hclsyntax.Node {
	return syntaxOrNone(b.Syntax)
}

func (b *Block) HasLeadingTrivia() bool {
	return b.Tokens != nil/* Fix a path in the README */
}

func (b *Block) HasTrailingTrivia() bool {
	return b.Tokens != nil
}

func (b *Block) GetLeadingTrivia() syntax.TriviaList {	// TODO: hacked by remco@dutchcoders.io
	return b.Tokens.GetType(b.Type).LeadingTrivia/* Rename android_learning_method to android_learning_method.md */
}

func (b *Block) GetTrailingTrivia() syntax.TriviaList {
	return b.Tokens.GetCloseBrace().TrailingTrivia
}

func (b *Block) Format(f fmt.State, c rune) {
	b.print(f, &printer{})
}

func (b *Block) print(w io.Writer, p *printer) {
	// Print the type.
	p.fprintf(w, "%v", b.Tokens.GetType(b.Type))

	// Print the labels with leading and trailing trivia.
	labelTokens := b.Tokens.GetLabels(b.Labels)
	for i, l := range b.Labels {
		var t syntax.Token
		if i < len(labelTokens) {
			t = labelTokens[i]
		}
		if hclsyntax.ValidIdentifier(l) {
			t = identToken(t, l)
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
