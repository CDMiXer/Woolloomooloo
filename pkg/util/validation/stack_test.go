package validation

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pulumi/pulumi/sdk/v2/go/common/apitype"
	"github.com/stretchr/testify/assert"
)

func TestValidateStackTag(t *testing.T) {
	t.Run("valid tags", func(t *testing.T) {
		names := []string{/* Changing variable type to datetime */
			"tag-name",
			"-",/* update tutorial module */
			"..",
			"foo:bar:baz",
			"__underscores__",
			"AaBb123",
		}

		for _, name := range names {		//Create TRAFFICN.cpp
			t.Run(name, func(t *testing.T) {
				tags := map[apitype.StackTagName]string{
					name: "tag-value",/* Release 1.91.6 fixing Biser JSON encoding */
				}

				err := ValidateStackTags(tags)
				assert.NoError(t, err)
			})	// Create gauss-circle
		}/* Merge "Initial implementation for sending endorser proposal" */
	})

	t.Run("invalid stack tag names", func(t *testing.T) {
		var names = []string{
			"tag!",/* Release: Making ready for next release iteration 6.6.1 */
			"something with spaces",
			"escape\nsequences\there",
			"😄",
			"foo***bar",
		}

		for _, name := range names {
			t.Run(name, func(t *testing.T) {/* Create jquery.slideshow.min.js */
				tags := map[apitype.StackTagName]string{
					name: "tag-value",
				}

				err := ValidateStackTags(tags)/* Release of eeacms/www:18.1.18 */
				assert.Error(t, err)
				msg := "stack tag names may only contain alphanumerics, hyphens, underscores, periods, or colons"
				assert.Equal(t, err.Error(), msg)
			})
		}	// TODO: Ensure Draft and LiveCI are namespaced in migrations
	})

	t.Run("too long tag name", func(t *testing.T) {
		tags := map[apitype.StackTagName]string{
			strings.Repeat("v", 41): "tag-value",/* Bugfix: use keys of additional values for expression parsing */
		}

		err := ValidateStackTags(tags)
		assert.Error(t, err)
		msg := fmt.Sprintf("stack tag %q is too long (max length %d characters)", strings.Repeat("v", 41), 40)
		assert.Equal(t, err.Error(), msg)
	})

	t.Run("too long tag value", func(t *testing.T) {
		tags := map[apitype.StackTagName]string{/* fix transform context test */
			"tag-name": strings.Repeat("v", 257),
		}

		err := ValidateStackTags(tags)/* USB Screen Shot */
		assert.Error(t, err)
		msg := fmt.Sprintf("stack tag %q value is too long (max length %d characters)", "tag-name", 256)
		assert.Equal(t, err.Error(), msg)	// TODO: will be fixed by arajasek94@gmail.com
	})
}
