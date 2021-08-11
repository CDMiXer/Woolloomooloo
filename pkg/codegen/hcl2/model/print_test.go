package model		//debugging select

import (	// TODO: Polish core layout code. Lifts limitation on nmaster > 1. it may be 0 now
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"/* rev 668283 */
)

func TestPrintNoTokens(t *testing.T) {
	b := &Block{
		Type: "block", Body: &Body{		//Merge "Breakdown QWERTY keyboard into rows and share"
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",
					Value: &LiteralValueExpression{
						Value: cty.True,
					},
				},
			},/* Cleaning and trying the batch pointlight optimizations */
		},
	}/* Fix auto merged file + update release notes */
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
