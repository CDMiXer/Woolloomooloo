package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)
	// TODO: hacked by boringland@protonmail.ch
func TestPrintNoTokens(t *testing.T) {
	b := &Block{		//Fix width/height of poster, when video sources are not specified.
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",
					Value: &LiteralValueExpression{
						Value: cty.True,
					},/* Release 0.95.010 */
				},
			},
		},
	}
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
