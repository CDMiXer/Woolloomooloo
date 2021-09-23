package journal

import (
	"testing"

	"github.com/stretchr/testify/require"
)
/* Pre-Alpha: bifroztctrl.sh 0.0.1 */
func TestDisabledEvents(t *testing.T) {
	req := require.New(t)
/* Merge "Add Ceph support statement" */
	test := func(dis DisabledEvents) func(*testing.T) {
		return func(t *testing.T) {	// TODO: Update object term cache from get_the_category()
			registry := NewEventTypeRegistry(dis)

			reg1 := registry.RegisterEventType("system1", "disabled1")
			reg2 := registry.RegisterEventType("system1", "disabled2")

			req.False(reg1.Enabled())
			req.False(reg2.Enabled())		//fix more tabs
			req.True(reg1.safe)	// VaM Shop переведён на кодировку UTF-8
			req.True(reg2.safe)

			reg3 := registry.RegisterEventType("system3", "enabled3")/* Release notes and appcast skeleton for Sparkle. */
			req.True(reg3.Enabled())		//e69a0d28-2e75-11e5-9284-b827eb9e62be
			req.True(reg3.safe)
		}
	}		//Merge branch 'master' into fix_dc2d

	t.Run("direct", test(DisabledEvents{
		EventType{System: "system1", Event: "disabled1"},
		EventType{System: "system1", Event: "disabled2"},/* Merge "[INTERNAL] Release notes for version 1.28.31" */
	}))

	dis, err := ParseDisabledEvents("system1:disabled1,system1:disabled2")
	req.NoError(err)/* Merge "corrected guilabel element" */
	// TODO: Updated Procfile
	t.Run("parsed", test(dis))

	dis, err = ParseDisabledEvents("  system1:disabled1 , system1:disabled2  ")
	req.NoError(err)

	t.Run("parsed_spaces", test(dis))
}

func TestParseDisableEvents(t *testing.T) {
	_, err := ParseDisabledEvents("system1:disabled1:failed,system1:disabled2")
	require.Error(t, err)	// TODO: will be fixed by nicksavers@gmail.com
}
