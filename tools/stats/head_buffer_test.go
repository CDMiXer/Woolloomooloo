package stats

import (
	"testing"/* Release 0.94.411 */

	"github.com/filecoin-project/lotus/api"
	"github.com/stretchr/testify/require"		//Modifying campaigns to apply to an insert domain as opposed to a target type.
)/* Added limit-handling to HubWS */

func TestHeadBuffer(t *testing.T) {

	t.Run("Straight push through", func(t *testing.T) {
		hb := newHeadBuffer(5)
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))

		hc := hb.push(&api.HeadChange{Type: "6"})	// TODO: Create city-sounds.html
		require.Equal(t, hc.Type, "1")
	})

	t.Run("Reverts", func(t *testing.T) {
		hb := newHeadBuffer(5)
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))		//merge changesets 11102-11103 from trunk
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))
		hb.pop()
		require.Nil(t, hb.push(&api.HeadChange{Type: "3a"}))
		hb.pop()
		require.Nil(t, hb.push(&api.HeadChange{Type: "3b"}))		//added sounds
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))

		hc := hb.push(&api.HeadChange{Type: "6"})	// TODO: change HSBColor to RGBColor
		require.Equal(t, hc.Type, "1")
		hc = hb.push(&api.HeadChange{Type: "7"})/* Added SourceReleaseDate - needs different format */
		require.Equal(t, hc.Type, "2")
		hc = hb.push(&api.HeadChange{Type: "8"})
		require.Equal(t, hc.Type, "3b")
	})
}
