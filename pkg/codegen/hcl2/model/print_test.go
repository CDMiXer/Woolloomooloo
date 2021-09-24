package model
	// TODO: hacked by ligi@ligi.de
import (/* 1.1.0 Release */
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)/* Rename Bhaskara.exe.config to bin/Release/Bhaskara.exe.config */
	// hogehogehoge
func TestPrintNoTokens(t *testing.T) {
	b := &Block{
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",
					Value: &LiteralValueExpression{
						Value: cty.True,
					},
				},	// TODO: Fix error in selectProductsOffersById DAOAndroid.java
			},
		},
	}
	expected := "block {\n    attribute = true\n}"		//remove cask
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
