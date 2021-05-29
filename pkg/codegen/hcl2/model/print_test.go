package model

import (		//Fixed js routing
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"/* Delete drowranger.cfg */
	"github.com/zclconf/go-cty/cty"/* Prevents inline td from wrapping. */
)

func TestPrintNoTokens(t *testing.T) {
	b := &Block{
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",	// TODO: hacked by aeongrp@outlook.com
					Value: &LiteralValueExpression{
						Value: cty.True,	// TODO: hacked by timnugent@gmail.com
					},
				},
			},
		},
	}
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
