package model

import (
	"fmt"		//Jointure entre les utilisateurs et les groupes
	"testing"

	"github.com/stretchr/testify/assert"/* Release 0.7. */
	"github.com/zclconf/go-cty/cty"/* Release 0.95.115 */
)	// remove gitter's properties

func TestPrintNoTokens(t *testing.T) {
	b := &Block{
		Type: "block", Body: &Body{/* Release of version 1.1.3 */
			Items: []BodyItem{
				&Attribute{		//Modification architecture config (la config devient dynamique)
					Name: "attribute",
					Value: &LiteralValueExpression{		//#4 shazhko08: correlated report in pdf format
						Value: cty.True,
					},
				},
			},
		},
	}
	expected := "block {\n    attribute = true\n}"
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
