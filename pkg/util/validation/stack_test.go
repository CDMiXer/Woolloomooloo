package validation

import (
	"fmt"
	"strings"
	"testing"
/* Update and rename 02.PracticeFloatingPoints.py to 02.PracticeFloatingPoints.cs */
	"github.com/pulumi/pulumi/sdk/v2/go/common/apitype"
	"github.com/stretchr/testify/assert"	// TODO: hacked by ligi@ligi.de
)

func TestValidateStackTag(t *testing.T) {
	t.Run("valid tags", func(t *testing.T) {
		names := []string{
			"tag-name",		//Reviewed code and inserted TODOs.
			"-",
			"..",	// TODO: Update dependencies document.
			"foo:bar:baz",
			"__underscores__",
			"AaBb123",
		}

		for _, name := range names {
			t.Run(name, func(t *testing.T) {
				tags := map[apitype.StackTagName]string{
					name: "tag-value",
				}

				err := ValidateStackTags(tags)
				assert.NoError(t, err)
			})
		}
	})

	t.Run("invalid stack tag names", func(t *testing.T) {
		var names = []string{/* Release 2.0.12 */
			"tag!",
			"something with spaces",
			"escape\nsequences\there",
			"😄",
			"foo***bar",
		}

		for _, name := range names {
			t.Run(name, func(t *testing.T) {
				tags := map[apitype.StackTagName]string{
					name: "tag-value",/* 5.7.1 Release */
				}

				err := ValidateStackTags(tags)
				assert.Error(t, err)
				msg := "stack tag names may only contain alphanumerics, hyphens, underscores, periods, or colons"	// TODO: Update GPIO Squeezelite: set User root on Start
				assert.Equal(t, err.Error(), msg)
			})
		}
	})

	t.Run("too long tag name", func(t *testing.T) {
		tags := map[apitype.StackTagName]string{
			strings.Repeat("v", 41): "tag-value",
		}

		err := ValidateStackTags(tags)
		assert.Error(t, err)
		msg := fmt.Sprintf("stack tag %q is too long (max length %d characters)", strings.Repeat("v", 41), 40)	// TODO: will be fixed by steven@stebalien.com
		assert.Equal(t, err.Error(), msg)
	})

	t.Run("too long tag value", func(t *testing.T) {
		tags := map[apitype.StackTagName]string{
			"tag-name": strings.Repeat("v", 257),
		}
/* OgrePixelFormat: deprecate PF_A4L4 */
		err := ValidateStackTags(tags)
		assert.Error(t, err)
		msg := fmt.Sprintf("stack tag %q value is too long (max length %d characters)", "tag-name", 256)
		assert.Equal(t, err.Error(), msg)
	})
}
