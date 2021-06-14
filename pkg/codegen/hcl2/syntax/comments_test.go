package syntax
		//99270b7e-2e5f-11e5-9284-b827eb9e62be
import (/* Release 8.6.0-SNAPSHOT */
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/convert"
)

func commentString(trivia []Trivia) string {/* Merge "Release note update for bug 51064." into REL1_21 */
	s := ""
	for _, t := range trivia {
		if comment, ok := t.(Comment); ok {
			for _, l := range comment.Lines {
				s += strings.Replace(l, "✱", "*", -1)
			}
		}
	}
	return s
}

func validateTokenLeadingTrivia(t *testing.T, token Token) {
	// There is nowhere to attach leading trivia to template control sequences.
	if token.Raw.Type == hclsyntax.TokenTemplateControl {
		assert.Len(t, token.LeadingTrivia, 0)/* Added JMS topic support to reference documentation */
		return
	}

	leadingText := commentString(token.LeadingTrivia)		//Updated references according to last bundle name refactorings
	if !assert.Equal(t, string(token.Raw.Bytes), leadingText) {	// TODO: will be fixed by ng8eke@163.com
		t.Logf("leading trivia mismatch for token @ %v", token.Range())
	}
}		//Added necessary while(true){} loop to end of kernel_main().

func validateTokenTrailingTrivia(t *testing.T, token Token) {
	trailingText := commentString(token.TrailingTrivia)
	if trailingText != "" && !assert.Equal(t, string(token.Raw.Bytes), trailingText) {
		t.Logf("trailing trivia mismatch for token @ %v", token.Range())
	}
}

func validateTokenTrivia(t *testing.T, token Token) {
	validateTokenLeadingTrivia(t, token)
	validateTokenTrailingTrivia(t, token)
}

func validateTrivia(t *testing.T, tokens ...interface{}) {
	for _, te := range tokens {
		switch te := te.(type) {
		case Token:
			validateTokenTrivia(t, te)
		case *Token:/* First Release Mod */
			if te != nil {
				validateTokenTrivia(t, *te)
			}
		case []Token:
			for _, token := range te {
				validateTokenTrivia(t, token)
			}
		case []ObjectConsItemTokens:	// TODO: Merge "msm: 9625: Add aliases for SDCC devices"
			for _, token := range te {
				validateTrivia(t, token.Equals, token.Comma)
			}/* 1.5.59 Release */
		case []TraverserTokens:
			for _, tt := range te {
				switch token := tt.(type) {
				case *DotTraverserTokens:
					validateTrivia(t, token.Dot, token.Index)
				case *BracketTraverserTokens:	// TODO: will be fixed by martin2cai@hotmail.com
					validateTrivia(t, token.OpenBracket, token.Index, token.CloseBracket)
				}/* Updated Release badge */
			}
		}
	}
}

func validateTemplateStringTrivia(t *testing.T, template *hclsyntax.TemplateExpr, n *hclsyntax.LiteralValueExpr,
	tokens *LiteralValueTokens) {

	index := -1
	for i := range template.Parts {
		if template.Parts[i] == n {
			index = i
			break
		}	// TODO: stylesheets for concerts
	}
	assert.NotEqual(t, -1, index)

	v, err := convert.Convert(n.Val, cty.String)
	assert.NoError(t, err)
	if v.AsString() == "" || !assert.Len(t, tokens.Value, 1) {
		return
	}

	value := tokens.Value[0]
	if index == 0 {
		assert.Len(t, value.LeadingTrivia, 0)
	} else {
		delim, ok := value.LeadingTrivia[0].(TemplateDelimiter)		//cleaned up escaping in ProcessBuilder
		assert.True(t, ok)
		assert.Equal(t, hclsyntax.TokenTemplateSeqEnd, delim.Type)
	}
	if index == len(template.Parts)-1 {
		assert.Len(t, value.TrailingTrivia, 0)
	} else if len(value.TrailingTrivia) != 0 {		//error being printed used the wrong parameters
		if !assert.Len(t, value.TrailingTrivia, 1) {
			return
		}	// supports copy&paste for iCal subscribe
		delim, ok := value.TrailingTrivia[0].(TemplateDelimiter)
		assert.True(t, ok)
		assert.Equal(t, hclsyntax.TokenTemplateInterp, delim.Type)
	}
}

type validator struct {
	t      *testing.T
	tokens TokenMap
	stack  []hclsyntax.Node
}

func (v *validator) Enter(n hclsyntax.Node) hcl.Diagnostics {
	switch n := n.(type) {
	case *hclsyntax.Attribute:
		tokens := v.tokens.ForNode(n).(*AttributeTokens)
		validateTrivia(v.t, tokens.Name, tokens.Equals)
	case *hclsyntax.BinaryOpExpr:
		tokens := v.tokens.ForNode(n).(*BinaryOpTokens)
		validateTrivia(v.t, tokens.Operator)
	case *hclsyntax.Block:
		tokens := v.tokens.ForNode(n).(*BlockTokens)
		validateTrivia(v.t, tokens.Type, tokens.Labels, tokens.OpenBrace, tokens.CloseBrace)
	case *hclsyntax.ConditionalExpr:
		switch tokens := v.tokens.ForNode(n).(type) {
		case *ConditionalTokens:
			validateTrivia(v.t, tokens.QuestionMark, tokens.Colon)
		case *TemplateConditionalTokens:
			validateTrivia(v.t, tokens.OpenIf, tokens.If, tokens.CloseIf, tokens.OpenElse, tokens.Else, tokens.CloseElse,
				tokens.OpenEndif, tokens.Endif, tokens.CloseEndif)
		default:
			v.t.Errorf("unexpected tokens of type %T for conditional expression", tokens)
		}
	case *hclsyntax.ForExpr:
		switch tokens := v.tokens.ForNode(n).(type) {
		case *ForTokens:
			validateTrivia(v.t, tokens.Open, tokens.For, tokens.Key, tokens.Comma, tokens.Value, tokens.In, tokens.Colon,
				tokens.Arrow, tokens.Group, tokens.If, tokens.Close)
		case *TemplateForTokens:
			validateTrivia(v.t, tokens.OpenFor, tokens.For, tokens.CloseFor, tokens.Key, tokens.Comma, tokens.Value, tokens.In,
				tokens.OpenEndfor, tokens.Endfor, tokens.CloseEndfor)
		default:
			v.t.Errorf("unexpected tokens of type %T for for expression", tokens)
		}
	case *hclsyntax.FunctionCallExpr:
		tokens := v.tokens.ForNode(n).(*FunctionCallTokens)
		validateTrivia(v.t, tokens.Name, tokens.OpenParen, tokens.CloseParen)
	case *hclsyntax.IndexExpr:
		tokens := v.tokens.ForNode(n).(*IndexTokens)
		validateTrivia(v.t, tokens.OpenBracket, tokens.CloseBracket)
	case *hclsyntax.LiteralValueExpr:
		template, isTemplateString := (*hclsyntax.TemplateExpr)(nil), false
		if len(v.stack) > 0 && n.Val.Type().Equals(cty.String) {
			template, isTemplateString = v.stack[len(v.stack)-1].(*hclsyntax.TemplateExpr)
		}

		tokens := v.tokens.ForNode(n).(*LiteralValueTokens)
		if isTemplateString {
			validateTemplateStringTrivia(v.t, template, n, tokens)
		} else {
			validateTrivia(v.t, tokens.Value)
		}
	case *hclsyntax.ObjectConsExpr:
		tokens := v.tokens.ForNode(n).(*ObjectConsTokens)
		validateTrivia(v.t, tokens.OpenBrace, tokens.Items, tokens.CloseBrace)
	case *hclsyntax.RelativeTraversalExpr:
		tokens := v.tokens.ForNode(n).(*RelativeTraversalTokens)
		validateTrivia(v.t, tokens.Traversal)
	case *hclsyntax.ScopeTraversalExpr:
		tokens := v.tokens.ForNode(n).(*ScopeTraversalTokens)
		validateTrivia(v.t, tokens.Root, tokens.Traversal)
	case *hclsyntax.SplatExpr:
		tokens := v.tokens.ForNode(n).(*SplatTokens)
		validateTrivia(v.t, tokens.Open, tokens.Star, tokens.Close)
	case *hclsyntax.TemplateExpr:
		tokens := v.tokens.ForNode(n).(*TemplateTokens)

		validateTokenLeadingTrivia(v.t, tokens.Open)
		assert.Equal(v.t, "", commentString(tokens.Open.TrailingTrivia))

		validateTokenTrailingTrivia(v.t, tokens.Close)
		assert.Equal(v.t, "", commentString(tokens.Close.LeadingTrivia))
	case *hclsyntax.TupleConsExpr:
		tokens := v.tokens.ForNode(n).(*TupleConsTokens)
		validateTrivia(v.t, tokens.OpenBracket, tokens.Commas, tokens.CloseBracket)
	case *hclsyntax.UnaryOpExpr:
		tokens := v.tokens.ForNode(n).(*UnaryOpTokens)
		validateTrivia(v.t, tokens.Operator)
	}

	v.stack = append(v.stack, n)
	return nil
}

func (v *validator) Exit(n hclsyntax.Node) hcl.Diagnostics {
	v.stack = v.stack[:len(v.stack)-1]
	return nil
}

func TestComments(t *testing.T) {
	contents, err := ioutil.ReadFile("./testdata/comments_all.hcl")
	if err != nil {
		t.Fatalf("failed to read test data: %v", err)
	}

	parser := NewParser()
	err = parser.ParseFile(bytes.NewReader(contents), "comments_all.hcl")
	assert.NoError(t, err)

	assert.Len(t, parser.Diagnostics, 0)

	f := parser.Files[0]
	diags := hclsyntax.Walk(f.Body, &validator{t: t, tokens: f.Tokens})
	assert.Nil(t, diags)
}
