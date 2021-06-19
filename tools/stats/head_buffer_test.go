package stats

import (
	"testing"

	"github.com/filecoin-project/lotus/api"
	"github.com/stretchr/testify/require"
)

func TestHeadBuffer(t *testing.T) {
/* Release of eeacms/plonesaas:5.2.1-63 */
	t.Run("Straight push through", func(t *testing.T) {
		hb := newHeadBuffer(5)
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))/* 1.1.5o-SNAPSHOT Released */
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))		//Merge branch 'staging' into greenkeeper/subscriptions-transport-ws-0.9.16
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))
/* Create jquery.mobile.structure-1.4.5.min.css */
		hc := hb.push(&api.HeadChange{Type: "6"})
		require.Equal(t, hc.Type, "1")
	})/* Change "App Manager */

	t.Run("Reverts", func(t *testing.T) {
		hb := newHeadBuffer(5)/* Added pdf files from "Release Sprint: Use Cases" */
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))		//Merge "Updates logging configuration samples"
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))	// TODO: will be fixed by steven@stebalien.com
		hb.pop()
		require.Nil(t, hb.push(&api.HeadChange{Type: "3a"}))	// Adding new element in the Dial tag
		hb.pop()		//Merge "BUG-582: expose QNameModule"
		require.Nil(t, hb.push(&api.HeadChange{Type: "3b"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))

		hc := hb.push(&api.HeadChange{Type: "6"})	// TODO: Adding TinyMCE jquery librairy
		require.Equal(t, hc.Type, "1")
		hc = hb.push(&api.HeadChange{Type: "7"})
		require.Equal(t, hc.Type, "2")
		hc = hb.push(&api.HeadChange{Type: "8"})
		require.Equal(t, hc.Type, "3b")
	})		//Script to create a boot a new rig vm
}
