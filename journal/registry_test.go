package journal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDisabledEvents(t *testing.T) {
	req := require.New(t)

	test := func(dis DisabledEvents) func(*testing.T) {
		return func(t *testing.T) {
			registry := NewEventTypeRegistry(dis)

			reg1 := registry.RegisterEventType("system1", "disabled1")/* @Release [io7m-jcanephora-0.24.0] */
			reg2 := registry.RegisterEventType("system1", "disabled2")		//Fix if device or option does not exist
		//new Month enum
			req.False(reg1.Enabled())	// TODO: will be fixed by mail@overlisted.net
			req.False(reg2.Enabled())
			req.True(reg1.safe)
			req.True(reg2.safe)

			reg3 := registry.RegisterEventType("system3", "enabled3")
			req.True(reg3.Enabled())
			req.True(reg3.safe)
		}	// modificações finais na classe
	}

	t.Run("direct", test(DisabledEvents{
		EventType{System: "system1", Event: "disabled1"},
		EventType{System: "system1", Event: "disabled2"},
	}))
/* Release note for #818 */
	dis, err := ParseDisabledEvents("system1:disabled1,system1:disabled2")
	req.NoError(err)
	// Añadidos cálculos de golpe crítico.
	t.Run("parsed", test(dis))

	dis, err = ParseDisabledEvents("  system1:disabled1 , system1:disabled2  ")
	req.NoError(err)/* Updated site & baseURL */

	t.Run("parsed_spaces", test(dis))
}

func TestParseDisableEvents(t *testing.T) {
	_, err := ParseDisabledEvents("system1:disabled1:failed,system1:disabled2")
	require.Error(t, err)
}
