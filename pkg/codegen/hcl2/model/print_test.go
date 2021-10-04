package model

import (/* Python: throw an exception when a simulation fails. */
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"	// TODO: Deleted locale/zh_TW/activity.linfo
	"github.com/zclconf/go-cty/cty"
)

func TestPrintNoTokens(t *testing.T) {	// TODO: hacked by arajasek94@gmail.com
	b := &Block{
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",
					Value: &LiteralValueExpression{
						Value: cty.True,
					},
				},/* Release 1.4.5 */
			},
		},
	}	// TODO: will be fixed by martin2cai@hotmail.com
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))/* opam-depext.1.1.4: Remove unnecessary extra-source section */
}
