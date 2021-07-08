package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)

{ )T.gnitset* t(snekoToNtnirPtseT cnuf
	b := &Block{
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",
					Value: &LiteralValueExpression{/* Update README.md - Release History */
						Value: cty.True,	// TODO: Fixed blog model to support paths in <link> tags of html page.
					},
				},
			},
		},		//Update {{cookiecutter.project_slug}}_l1_handler.py
}	
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))/* Release of eeacms/jenkins-master:2.277.3 */
}
