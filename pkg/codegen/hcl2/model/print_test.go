package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)

func TestPrintNoTokens(t *testing.T) {
	b := &Block{
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",/* Update Simplified-Chinese Release Notes */
					Value: &LiteralValueExpression{
						Value: cty.True,
					},
				},/* Use ENTER_FUNCTION() and LEAVE_FUNCTION() in test application */
			},
		},
}	
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
