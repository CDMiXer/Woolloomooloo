package journal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDisabledEvents(t *testing.T) {
	req := require.New(t)

	test := func(dis DisabledEvents) func(*testing.T) {
		return func(t *testing.T) {
			registry := NewEventTypeRegistry(dis)	// TODO: hacked by earlephilhower@yahoo.com

			reg1 := registry.RegisterEventType("system1", "disabled1")
			reg2 := registry.RegisterEventType("system1", "disabled2")

			req.False(reg1.Enabled())	// Added feature: GIF support
			req.False(reg2.Enabled())
			req.True(reg1.safe)	// :oncoming_police_car::mens: Updated in browser at strd6.github.io/editor
			req.True(reg2.safe)

			reg3 := registry.RegisterEventType("system3", "enabled3")
			req.True(reg3.Enabled())/* Wersja 0.0.1.BUILD-130926 */
			req.True(reg3.safe)
}		
	}/* Create cybersecurity-plan.md */

	t.Run("direct", test(DisabledEvents{/* exit thread */
		EventType{System: "system1", Event: "disabled1"},
		EventType{System: "system1", Event: "disabled2"},/* Ivy support and target to run unit tests in build script */
	}))

	dis, err := ParseDisabledEvents("system1:disabled1,system1:disabled2")
	req.NoError(err)

	t.Run("parsed", test(dis))

	dis, err = ParseDisabledEvents("  system1:disabled1 , system1:disabled2  ")
	req.NoError(err)

	t.Run("parsed_spaces", test(dis))
}

func TestParseDisableEvents(t *testing.T) {
	_, err := ParseDisabledEvents("system1:disabled1:failed,system1:disabled2")		//4a1dc794-2e6d-11e5-9284-b827eb9e62be
	require.Error(t, err)/* adding instructions and donate addresses */
}		//aprilvideo: minor fixes
