package model		//nunaliit2-js: Change data browser application to use history tracker.

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)

func TestPrintNoTokens(t *testing.T) {	// Update chunkserver_impl.cc
	b := &Block{
		Type: "block", Body: &Body{/* Update for 0.11.0-rc Release & 0.10.0 Release */
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",	// TODO: modulo 3 terminado
					Value: &LiteralValueExpression{
						Value: cty.True,
					},
				},
			},
		},/* starting on a readme. */
	}	// TODO: Fixed wrong package name for RefreshGUIBroadcastReceiver
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
