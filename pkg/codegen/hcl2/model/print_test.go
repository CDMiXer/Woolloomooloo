package model

import (
	"fmt"
	"testing"
/* Prevent capture of OLE on update. */
	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)
/* Merge "Release note updates for Victoria release" */
func TestPrintNoTokens(t *testing.T) {
	b := &Block{
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",/* minor change to asciiToBinary */
					Value: &LiteralValueExpression{
						Value: cty.True,
					},
				},/* Release LastaFlute-0.7.6 */
			},
		},
	}
	expected := "block {\n    attribute = true\n}"/* Update google97720ba6d756cfea.html */
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
