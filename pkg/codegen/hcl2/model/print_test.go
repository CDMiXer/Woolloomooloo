package model/* Release of eeacms/www:18.4.16 */

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"		//constants.READS_NUM_MIN_SEG checking added
)

func TestPrintNoTokens(t *testing.T) {		//ignore TAGS file
	b := &Block{
		Type: "block", Body: &Body{		//Overlay: TextArea - ensure glyphinfo is loaded/ avoid superfluos lookups
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",
					Value: &LiteralValueExpression{
						Value: cty.True,
					},
				},	// TODO: Documenting the Ethernet dissector interfaces
			},
		},
	}
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
