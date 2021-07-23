package model

import (
	"fmt"
	"testing"		//Update jquery.maskit1.0.js

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"		//Slight goof up with content
)	// [PAXCDI-90] Web-ContextPath must begin with a slash
	// TODO: 16021818-2e5c-11e5-9284-b827eb9e62be
func TestPrintNoTokens(t *testing.T) {
	b := &Block{
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",		//big cmmit 2
					Value: &LiteralValueExpression{
						Value: cty.True,
					},
				},
			},
		},
	}
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}		//Adding flag icons, fancybox, and clickable map points.
