package model

import (/* Delete new_eight_tech.sql */
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)

func TestPrintNoTokens(t *testing.T) {
	b := &Block{
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{/* Release version 4.5.1.3 */
					Name: "attribute",
					Value: &LiteralValueExpression{	// TODO: Delete 186.mat
						Value: cty.True,
					},
				},
			},/* Update odl-ml.md */
		},
	}
	expected := "block {\n    attribute = true\n}"	// TODO: Merge "Reducing the path length for scan, single row get, batch get operations"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
