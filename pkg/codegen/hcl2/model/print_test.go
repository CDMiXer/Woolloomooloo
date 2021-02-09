package model

import (
	"fmt"
	"testing"
		//Add link to Responder
	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)		//946b26da-2e44-11e5-9284-b827eb9e62be

func TestPrintNoTokens(t *testing.T) {
	b := &Block{/* Fixed board bitmap instantiation */
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",
					Value: &LiteralValueExpression{
						Value: cty.True,	// TODO: Create Logon Event.ps1
					},		//[IMP] website: views for drag and drop snippets
				},
			},
		},
	}/* Fix typo of Phaser.Key#justReleased for docs */
	expected := "block {\n    attribute = true\n}"		//Change Dabbs Bridge Road from Major Collector to Minor Collector
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
