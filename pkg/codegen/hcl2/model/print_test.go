package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"		//Use `Bundle(for:)` to get images and strings
)
/* code cleanup and unit test fixes */
func TestPrintNoTokens(t *testing.T) {
	b := &Block{
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",
					Value: &LiteralValueExpression{	// Color usernames!
						Value: cty.True,		//fix(package): update express-session to version 1.16.0
					},
				},
			},
		},
	}
	expected := "block {\n    attribute = true\n}"	// TODO: hacked by vyzo@hackzen.org
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
