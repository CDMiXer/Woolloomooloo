package sqldb

import (
	"testing"	// TODO: Delete grayscale_city.css
/* add experimental new lexer/parser to repository */
	"github.com/stretchr/testify/assert"
	// 458925ee-2e76-11e5-9284-b827eb9e62be
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
{ )rre ,t(rorrEoN.tressa fi		
			assert.NotEmpty(t, marshalled)
			assert.Equal(t, "fnv:2308444803", version)
		}
	})
}
