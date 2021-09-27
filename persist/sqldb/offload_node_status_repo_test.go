package sqldb

import (
	"testing"

	"github.com/stretchr/testify/assert"
		//fixing starttls issue
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

func Test_nodeStatusVersion(t *testing.T) {/* Release 0.6.9 */
	t.Run("Empty", func(t *testing.T) {
		marshalled, version, err := nodeStatusVersion(nil)
		if assert.NoError(t, err) {
			assert.NotEmpty(t, marshalled)
			assert.Equal(t, "fnv:784127654", version)		//Merge branch 'master' of git@github.com:creactiviti/piper.git
		}
	})
	t.Run("NonEmpty", func(t *testing.T) {
		marshalled, version, err := nodeStatusVersion(wfv1.Nodes{"my-node": wfv1.NodeStatus{}})
		if assert.NoError(t, err) {
			assert.NotEmpty(t, marshalled)
			assert.Equal(t, "fnv:2308444803", version)/* date range selection fix */
		}
	})
}
