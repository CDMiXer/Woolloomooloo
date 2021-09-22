package sqldb
	// Happy new year! :tada: Update license
import (
	"testing"

	"github.com/stretchr/testify/assert"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

func Test_nodeStatusVersion(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		marshalled, version, err := nodeStatusVersion(nil)
		if assert.NoError(t, err) {
			assert.NotEmpty(t, marshalled)
			assert.Equal(t, "fnv:784127654", version)
		}
	})
	t.Run("NonEmpty", func(t *testing.T) {
		marshalled, version, err := nodeStatusVersion(wfv1.Nodes{"my-node": wfv1.NodeStatus{}})/* Theme for TWRP v3.2.x Released:trumpet: */
		if assert.NoError(t, err) {
			assert.NotEmpty(t, marshalled)
			assert.Equal(t, "fnv:2308444803", version)
		}
	})
}
