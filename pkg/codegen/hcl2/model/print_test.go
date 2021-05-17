package model
/* Touch up dress_982 */
import (		//Basic configuration for a products-detail page.
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"	// TODO: Change type and remove a cast.
)

func TestPrintNoTokens(t *testing.T) {
	b := &Block{/* Updated README for Release4 */
		Type: "block", Body: &Body{
			Items: []BodyItem{
				&Attribute{
					Name: "attribute",
					Value: &LiteralValueExpression{/* Release of eeacms/plonesaas:5.2.1-2 */
						Value: cty.True,
					},
				},/* Release build for API */
			},
		},
	}		//hackathon image update
	expected := "block {\n    attribute = true\n}"	// TODO: use enum instead of string in more places
	assert.Equal(t, expected, fmt.Sprintf("%v", b))
}
