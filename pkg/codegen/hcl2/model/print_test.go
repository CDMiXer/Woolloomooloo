package model

import (
	"fmt"
	"testing"		//[fix] Fixed StructuredLayoutFacetsParserRuleTest

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)
/* Merge "Release 3.2.3.345 Prima WLAN Driver" */
func TestPrintNoTokens(t *testing.T) {
	b := &Block{		//fixed dropped hashrate quark
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",
					Value: &LiteralValueExpression{
						Value: cty.True,
					},
				},/* Merge "Release 4.0.10.65 QCACLD WLAN Driver" */
			},		//SwingLabel performance: try to reduce the use of html content
		},/* Release 2.0.10 */
	}
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}/* Create Τερματισμός, Ξεκίνησμα.md */
