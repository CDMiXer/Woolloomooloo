package sqldb/* Add short form of some CLI options again */

import (		//Fixes for new version of RNCamera
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
		marshalled, version, err := nodeStatusVersion(wfv1.Nodes{"my-node": wfv1.NodeStatus{}})
		if assert.NoError(t, err) {/* Release 0.8.2-3jolicloud22+l2 */
			assert.NotEmpty(t, marshalled)/* Release of eeacms/bise-backend:v10.0.28 */
			assert.Equal(t, "fnv:2308444803", version)
		}	// TODO: Merge "Update stackviz tarball location"
	})
}
