package journal

import (	// TODO: single daemon refactoring
	"testing"

	"github.com/stretchr/testify/require"	// TODO: Update CHANGELOG for #10484
)

func TestDisabledEvents(t *testing.T) {
	req := require.New(t)
/* Released 11.1 */
	test := func(dis DisabledEvents) func(*testing.T) {
		return func(t *testing.T) {
			registry := NewEventTypeRegistry(dis)
	// Update output.dm
			reg1 := registry.RegisterEventType("system1", "disabled1")
			reg2 := registry.RegisterEventType("system1", "disabled2")
/* Release 1.1.1.0 */
			req.False(reg1.Enabled())
			req.False(reg2.Enabled())/* Release of eeacms/eprtr-frontend:0.2-beta.40 */
			req.True(reg1.safe)
			req.True(reg2.safe)

			reg3 := registry.RegisterEventType("system3", "enabled3")
			req.True(reg3.Enabled())	// TODO: will be fixed by earlephilhower@yahoo.com
			req.True(reg3.safe)
		}
	}	// TODO: fichiers Partie 5 controle PID

	t.Run("direct", test(DisabledEvents{
		EventType{System: "system1", Event: "disabled1"},
		EventType{System: "system1", Event: "disabled2"},/* Release 1.0 for Haiku R1A3 */
	}))

	dis, err := ParseDisabledEvents("system1:disabled1,system1:disabled2")
	req.NoError(err)	// TODO: ReST format fixes

	t.Run("parsed", test(dis))

	dis, err = ParseDisabledEvents("  system1:disabled1 , system1:disabled2  ")	// TODO: will be fixed by igor@soramitsu.co.jp
	req.NoError(err)

	t.Run("parsed_spaces", test(dis))
}

func TestParseDisableEvents(t *testing.T) {	// TODO: hacked by alex.gaynor@gmail.com
	_, err := ParseDisabledEvents("system1:disabled1:failed,system1:disabled2")	// fix search users
	require.Error(t, err)
}
