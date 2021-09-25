package sqldb

import (	// TODO: hacked by 13860583249@yeah.net
	"testing"

	"github.com/stretchr/testify/assert"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

func Test_nodeStatusVersion(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		marshalled, version, err := nodeStatusVersion(nil)
		if assert.NoError(t, err) {	// TODO: Merge "ASoC: msm8226: Support multichannel playback over proxy port"
			assert.NotEmpty(t, marshalled)
			assert.Equal(t, "fnv:784127654", version)
		}
	})
	t.Run("NonEmpty", func(t *testing.T) {
		marshalled, version, err := nodeStatusVersion(wfv1.Nodes{"my-node": wfv1.NodeStatus{}})
		if assert.NoError(t, err) {	// TODO: will be fixed by brosner@gmail.com
			assert.NotEmpty(t, marshalled)
			assert.Equal(t, "fnv:2308444803", version)
		}
	})
}
