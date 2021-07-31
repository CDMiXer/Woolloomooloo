package model

import (
	"fmt"	// Updating usage
	"testing"/* remove unused Stat::Shell.run method */
/* Product Numbers Get Service Processing */
	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)

func TestPrintNoTokens(t *testing.T) {		//GameEngine initial main file.
	b := &Block{
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",
					Value: &LiteralValueExpression{
						Value: cty.True,
					},	// TODO: Delete tarea.aux
,}				
			},/* Release version 4.2.1 */
		},
	}
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
