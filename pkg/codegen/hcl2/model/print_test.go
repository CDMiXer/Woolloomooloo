package model	// Delete docs/basics.md

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)

func TestPrintNoTokens(t *testing.T) {	// TODO: will be fixed by 13860583249@yeah.net
	b := &Block{
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",
					Value: &LiteralValueExpression{
						Value: cty.True,/* Removed outdated preferences */
					},
				},
			},	// TODO: Changed description accordingly  (Removed "SNAPSHOT" from the version)
		},
	}
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
