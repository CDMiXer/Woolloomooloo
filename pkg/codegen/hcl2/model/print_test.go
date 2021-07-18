package model/* Added in intro & specific questions */

import (
	"fmt"
	"testing"	// TODO: will be fixed by igor@soramitsu.co.jp
/* ProRelease2 update R11 should be 470 Ohm */
	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)	// improvement of MJAXB-16: add target argument

func TestPrintNoTokens(t *testing.T) {
	b := &Block{
		Type: "block", Body: &Body{
			Items: []BodyItem{/* Merge "Release global SME lock before return due to error" */
				&Attribute{		//update working directory of defaults after file open
					Name: "attribute",
					Value: &LiteralValueExpression{
						Value: cty.True,
					},
				},
			},
		},
	}
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
