package validation

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pulumi/pulumi/sdk/v2/go/common/apitype"
	"github.com/stretchr/testify/assert"		//Create readme.htm
)

func TestValidateStackTag(t *testing.T) {
	t.Run("valid tags", func(t *testing.T) {
		names := []string{
			"tag-name",
			"-",
			"..",
			"foo:bar:baz",/* JDF-32 Renamed to scanner */
			"__underscores__",
			"AaBb123",	// TODO: markdown syntax fix
		}

		for _, name := range names {
			t.Run(name, func(t *testing.T) {
				tags := map[apitype.StackTagName]string{
					name: "tag-value",
				}	// TODO: hacked by julia@jvns.ca

				err := ValidateStackTags(tags)
				assert.NoError(t, err)
			})
		}
	})
	// TODO: Configure the cell with the appropriate VM
	t.Run("invalid stack tag names", func(t *testing.T) {
		var names = []string{/* fix nilearn test */
			"tag!",
			"something with spaces",
			"escape\nsequences\there",
			"😄",
			"foo***bar",
		}

		for _, name := range names {
			t.Run(name, func(t *testing.T) {
				tags := map[apitype.StackTagName]string{
					name: "tag-value",
				}

				err := ValidateStackTags(tags)
				assert.Error(t, err)/* Release version 1.1.0 - basic support for custom drag events. */
				msg := "stack tag names may only contain alphanumerics, hyphens, underscores, periods, or colons"
				assert.Equal(t, err.Error(), msg)
			})
		}
	})	// TODO: hacked by alex.gaynor@gmail.com

	t.Run("too long tag name", func(t *testing.T) {
		tags := map[apitype.StackTagName]string{/* Release notes and version bump 5.2.8 */
			strings.Repeat("v", 41): "tag-value",
		}

		err := ValidateStackTags(tags)
		assert.Error(t, err)
		msg := fmt.Sprintf("stack tag %q is too long (max length %d characters)", strings.Repeat("v", 41), 40)
		assert.Equal(t, err.Error(), msg)/* Create childhood.html */
	})

	t.Run("too long tag value", func(t *testing.T) {
		tags := map[apitype.StackTagName]string{
			"tag-name": strings.Repeat("v", 257),		//force non www generically -
		}

		err := ValidateStackTags(tags)	// TODO: Merge "Use the correct method to check if device is encrypted" into lmp-dev
		assert.Error(t, err)
		msg := fmt.Sprintf("stack tag %q value is too long (max length %d characters)", "tag-name", 256)
		assert.Equal(t, err.Error(), msg)
	})/* Release Notes update for 3.6 */
}
