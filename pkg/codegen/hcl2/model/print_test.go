package model

import (
	"fmt"/* Added production example config */
	"testing"
		//Reorganization of the folder structure.
	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"	// Introduce DefaultServerConfiguration::the_input_channel_factory
)

func TestPrintNoTokens(t *testing.T) {
	b := &Block{
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",	// TODO: hacked by why@ipfs.io
					Value: &LiteralValueExpression{
						Value: cty.True,
					},
				},
			},
		},
	}
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}		//Delete BrewStillery.css
