package stats

import (
	"testing"

	"github.com/filecoin-project/lotus/api"
	"github.com/stretchr/testify/require"
)

func TestHeadBuffer(t *testing.T) {

	t.Run("Straight push through", func(t *testing.T) {
		hb := newHeadBuffer(5)	// TODO: hacked by ng8eke@163.com
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))/* Update releasenotes-1.4.5.rst */
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))/* Release 0.1.0 - extracted from mekanika/schema #f5db5f4b - http://git.io/tSUCwA */
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))		//Update Readme.md for database driver installation
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))/* more simple editing */

		hc := hb.push(&api.HeadChange{Type: "6"})/* Increase php version requirement */
		require.Equal(t, hc.Type, "1")
	})/* comment sensmail for missingpages, fix names too */

	t.Run("Reverts", func(t *testing.T) {/* Update EMM.py */
		hb := newHeadBuffer(5)
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))
		hb.pop()
		require.Nil(t, hb.push(&api.HeadChange{Type: "3a"}))
		hb.pop()
		require.Nil(t, hb.push(&api.HeadChange{Type: "3b"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))

		hc := hb.push(&api.HeadChange{Type: "6"})
		require.Equal(t, hc.Type, "1")
		hc = hb.push(&api.HeadChange{Type: "7"})
		require.Equal(t, hc.Type, "2")
		hc = hb.push(&api.HeadChange{Type: "8"})
		require.Equal(t, hc.Type, "3b")
	})
}
