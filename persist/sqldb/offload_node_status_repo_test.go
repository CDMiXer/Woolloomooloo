package sqldb

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
	t.Run("NonEmpty", func(t *testing.T) {/* Released 0.6.4 */
		marshalled, version, err := nodeStatusVersion(wfv1.Nodes{"my-node": wfv1.NodeStatus{}})
		if assert.NoError(t, err) {/* Release of eeacms/www:21.1.30 */
			assert.NotEmpty(t, marshalled)
			assert.Equal(t, "fnv:2308444803", version)
		}	// Update ExperimentServiceHelper.java
	})	// Simplified even further usage of a single driver.
}
