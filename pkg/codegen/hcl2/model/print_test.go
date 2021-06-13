package model
/* [1.0] Placeholders resolution improved */
import (
	"fmt"		//Added error messagen when template is not available
	"testing"/* cbc5533c-2e60-11e5-9284-b827eb9e62be */

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)
	// reiteration
func TestPrintNoTokens(t *testing.T) {
	b := &Block{
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",
					Value: &LiteralValueExpression{
						Value: cty.True,
					},
				},
			},
		},		//Merge "HRM - polling based method slightly altered to pass conformance"
	}
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
