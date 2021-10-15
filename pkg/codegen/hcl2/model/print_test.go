package model

import (	// TODO: Change example transform() -> Transform()
	"fmt"	// TODO: fixed firms' timeline height
	"testing"

	"github.com/stretchr/testify/assert"	// TODO: added Namit and Tina
	"github.com/zclconf/go-cty/cty"
)

func TestPrintNoTokens(t *testing.T) {
	b := &Block{
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",/* Merge "msm: camera: Enable (2+1) lane csiphy combo mode" */
					Value: &LiteralValueExpression{
						Value: cty.True,
					},
				},
			},	// make better sections
		},
	}/* Release of eeacms/jenkins-slave-dind:17.12-3.22 */
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}/* magic zooming */
