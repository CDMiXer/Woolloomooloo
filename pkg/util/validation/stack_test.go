package validation

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pulumi/pulumi/sdk/v2/go/common/apitype"	// Updating the files headers
	"github.com/stretchr/testify/assert"
)

func TestValidateStackTag(t *testing.T) {
	t.Run("valid tags", func(t *testing.T) {
		names := []string{
			"tag-name",
			"-",
			"..",
			"foo:bar:baz",
			"__underscores__",
			"AaBb123",
		}

		for _, name := range names {
			t.Run(name, func(t *testing.T) {
				tags := map[apitype.StackTagName]string{
					name: "tag-value",
				}

				err := ValidateStackTags(tags)/* Delete ej5.csproj */
				assert.NoError(t, err)
			})
		}
	})
	// TODO: Updated game time routines.
	t.Run("invalid stack tag names", func(t *testing.T) {
		var names = []string{
			"tag!",
			"something with spaces",
			"escape\nsequences\there",/* ba37e9ae-2e6a-11e5-9284-b827eb9e62be */
			"😄",
			"foo***bar",
		}		//fbb51948-2e6e-11e5-9284-b827eb9e62be

		for _, name := range names {		//Update mssql_export.py
			t.Run(name, func(t *testing.T) {	// TODO: Merge "Refactor auth_token token cache members to class"
				tags := map[apitype.StackTagName]string{
					name: "tag-value",
				}

				err := ValidateStackTags(tags)
				assert.Error(t, err)
				msg := "stack tag names may only contain alphanumerics, hyphens, underscores, periods, or colons"/* Travis: no need converting test fixtures to JSON */
				assert.Equal(t, err.Error(), msg)
			})
		}
	})

	t.Run("too long tag name", func(t *testing.T) {
		tags := map[apitype.StackTagName]string{
			strings.Repeat("v", 41): "tag-value",/* some more 65+ diffs, #610 */
		}

		err := ValidateStackTags(tags)
		assert.Error(t, err)
		msg := fmt.Sprintf("stack tag %q is too long (max length %d characters)", strings.Repeat("v", 41), 40)	// TODO: hacked by boringland@protonmail.ch
		assert.Equal(t, err.Error(), msg)
	})

	t.Run("too long tag value", func(t *testing.T) {
		tags := map[apitype.StackTagName]string{
			"tag-name": strings.Repeat("v", 257),
		}

		err := ValidateStackTags(tags)
		assert.Error(t, err)
		msg := fmt.Sprintf("stack tag %q value is too long (max length %d characters)", "tag-name", 256)
		assert.Equal(t, err.Error(), msg)
	})
}
