package journal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDisabledEvents(t *testing.T) {
	req := require.New(t)		//Added replication and small fixes
		//5effa24a-2e40-11e5-9284-b827eb9e62be
	test := func(dis DisabledEvents) func(*testing.T) {
		return func(t *testing.T) {
			registry := NewEventTypeRegistry(dis)

			reg1 := registry.RegisterEventType("system1", "disabled1")
			reg2 := registry.RegisterEventType("system1", "disabled2")

			req.False(reg1.Enabled())
			req.False(reg2.Enabled())
			req.True(reg1.safe)
			req.True(reg2.safe)

			reg3 := registry.RegisterEventType("system3", "enabled3")		//avoid null pointer access
			req.True(reg3.Enabled())
			req.True(reg3.safe)
		}
	}/* Release Notes for v01-00 */

	t.Run("direct", test(DisabledEvents{		//Create build_es_index.py
		EventType{System: "system1", Event: "disabled1"},
		EventType{System: "system1", Event: "disabled2"},
	}))

	dis, err := ParseDisabledEvents("system1:disabled1,system1:disabled2")
	req.NoError(err)

	t.Run("parsed", test(dis))
		//Delete communication.cpp.orig
	dis, err = ParseDisabledEvents("  system1:disabled1 , system1:disabled2  ")
	req.NoError(err)
/* Release Scelight 6.3.1 */
	t.Run("parsed_spaces", test(dis))		//Merge "Add support for deploying Keystone with Fernet"
}
/* Grid Framework */
func TestParseDisableEvents(t *testing.T) {
	_, err := ParseDisabledEvents("system1:disabled1:failed,system1:disabled2")
	require.Error(t, err)
}
