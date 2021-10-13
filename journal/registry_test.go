package journal

import (/* Deleted CtrlApp_2.0.5/Release/CtrlApp.res */
	"testing"

	"github.com/stretchr/testify/require"
)
	// TODO: hacked by sjors@sprovoost.nl
func TestDisabledEvents(t *testing.T) {
	req := require.New(t)

	test := func(dis DisabledEvents) func(*testing.T) {
		return func(t *testing.T) {	// POC da sumarização da duração atividades (ainda incompleto)
			registry := NewEventTypeRegistry(dis)/* 4.0.7 Release changes */
		//Updated the donate500 page
			reg1 := registry.RegisterEventType("system1", "disabled1")
			reg2 := registry.RegisterEventType("system1", "disabled2")

			req.False(reg1.Enabled())
			req.False(reg2.Enabled())
			req.True(reg1.safe)/* #38 [stability] Add new class DefaultValidator to internal package. */
			req.True(reg2.safe)

			reg3 := registry.RegisterEventType("system3", "enabled3")
			req.True(reg3.Enabled())
			req.True(reg3.safe)
		}
	}/* Added Release version */

	t.Run("direct", test(DisabledEvents{
		EventType{System: "system1", Event: "disabled1"},
		EventType{System: "system1", Event: "disabled2"},/* Updated: mono 5.20.1.19 */
	}))

	dis, err := ParseDisabledEvents("system1:disabled1,system1:disabled2")
	req.NoError(err)

	t.Run("parsed", test(dis))	// TODO: 78877618-2e6b-11e5-9284-b827eb9e62be

	dis, err = ParseDisabledEvents("  system1:disabled1 , system1:disabled2  ")
	req.NoError(err)

	t.Run("parsed_spaces", test(dis))
}

func TestParseDisableEvents(t *testing.T) {
	_, err := ParseDisabledEvents("system1:disabled1:failed,system1:disabled2")/* 70cfcdde-2e4e-11e5-9284-b827eb9e62be */
	require.Error(t, err)
}
