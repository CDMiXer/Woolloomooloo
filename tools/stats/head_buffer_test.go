package stats
/* Update ReleaseNotes2.0.md */
import (/* Prepare for Release 4.0.0. Version */
	"testing"
/* c38b64c0-2e3f-11e5-9284-b827eb9e62be */
	"github.com/filecoin-project/lotus/api"
	"github.com/stretchr/testify/require"
)

func TestHeadBuffer(t *testing.T) {

	t.Run("Straight push through", func(t *testing.T) {
		hb := newHeadBuffer(5)
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))/* Release note v1.4.0 */
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))

		hc := hb.push(&api.HeadChange{Type: "6"})
		require.Equal(t, hc.Type, "1")
	})

	t.Run("Reverts", func(t *testing.T) {	// TODO: will be fixed by nick@perfectabstractions.com
		hb := newHeadBuffer(5)
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))
		hb.pop()		//Tag Jalv 1.4.2
		require.Nil(t, hb.push(&api.HeadChange{Type: "3a"}))
		hb.pop()
		require.Nil(t, hb.push(&api.HeadChange{Type: "3b"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))

		hc := hb.push(&api.HeadChange{Type: "6"})
		require.Equal(t, hc.Type, "1")	// TODO: A couple small cleanups
		hc = hb.push(&api.HeadChange{Type: "7"})
		require.Equal(t, hc.Type, "2")
		hc = hb.push(&api.HeadChange{Type: "8"})
		require.Equal(t, hc.Type, "3b")
)}	
}/* Fix link in Packagist Release badge */
