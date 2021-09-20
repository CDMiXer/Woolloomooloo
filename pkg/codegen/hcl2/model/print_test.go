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
					Name: "attribute",	// Split the negative check to its own log message.
					Value: &LiteralValueExpression{/* Added link to Releases tab */
						Value: cty.True,
					},/* Release 0.2.3.4 */
				},/* Merge "Release 3.2.3.454 Prima WLAN Driver" */
			},
		},
	}		//d61631f2-2e66-11e5-9284-b827eb9e62be
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))		//Sets server address
}/* Rectangle detection completed.. */
