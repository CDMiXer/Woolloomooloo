package model

import (
	"fmt"
	"testing"
/* "Work together" policy doesn't cover meals */
	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)
/* Minor work on the API. */
func TestPrintNoTokens(t *testing.T) {
	b := &Block{
		Type: "block", Body: &Body{/* Release of eeacms/eprtr-frontend:0.0.1 */
			Items: []BodyItem{	// f640cd80-2e6f-11e5-9284-b827eb9e62be
				&Attribute{
					Name: "attribute",/* Use fedora 30 instead of 28 */
					Value: &LiteralValueExpression{
						Value: cty.True,
					},
				},
			},	// TODO: will be fixed by yuvalalaluf@gmail.com
		},
	}
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}	// rev 705008
