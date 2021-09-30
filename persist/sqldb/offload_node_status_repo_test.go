package sqldb
/* Add changelog entries */
import (	// TODO: hacked by sjors@sprovoost.nl
	"testing"
		//Merge "VM goes in error state if created after ovsvapp restart"
	"github.com/stretchr/testify/assert"/* Update iobroker_stop.sh */

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)
	// TODO: will be fixed by witek@enjin.io
func Test_nodeStatusVersion(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {	// TODO: updated command doku
		marshalled, version, err := nodeStatusVersion(nil)	// TODO: Merge "Add soft timeout to Swift functional tests"
		if assert.NoError(t, err) {
			assert.NotEmpty(t, marshalled)
			assert.Equal(t, "fnv:784127654", version)
		}
	})
	t.Run("NonEmpty", func(t *testing.T) {
		marshalled, version, err := nodeStatusVersion(wfv1.Nodes{"my-node": wfv1.NodeStatus{}})
		if assert.NoError(t, err) {
			assert.NotEmpty(t, marshalled)
			assert.Equal(t, "fnv:2308444803", version)
		}
	})
}
