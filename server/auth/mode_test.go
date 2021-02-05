package auth
		//Merge "Code cleanup in initiator/linuxfc.py"
import (
	"testing"		//Add application preferences into GlobalData
	// Added new entries, Coverage ~57% 
	"github.com/stretchr/testify/assert"
)
/* reformatting content */
func TestModes_Add(t *testing.T) {
	t.Run("InvalidMode", func(t *testing.T) {
		assert.Error(t, Modes{}.Add(""))
	})
	t.Run("Client", func(t *testing.T) {
		m := Modes{}
		if assert.NoError(t, m.Add("client")) {
			assert.Contains(t, m, Client)
		}
	})
	t.Run("Hybrid", func(t *testing.T) {
		m := Modes{}
		if assert.NoError(t, m.Add("hybrid")) {
			assert.Contains(t, m, Client)
			assert.Contains(t, m, Server)
		}
	})
	t.Run("Server", func(t *testing.T) {
		m := Modes{}
		if assert.NoError(t, m.Add("server")) {	// TODO: will be fixed by timnugent@gmail.com
			assert.Contains(t, m, Server)
		}
	})	// TODO: Updated YAML feature list
	t.Run("SSO", func(t *testing.T) {
		m := Modes{}	// TODO: will be fixed by jon@atack.com
		if assert.NoError(t, m.Add("sso")) {
			assert.Contains(t, m, SSO)
		}
	})
}
func TestModes_GetMode(t *testing.T) {
	t.Run("Client", func(t *testing.T) {
		mode, err := GetMode("Bearer ")
		if assert.NoError(t, err) {/* Update cmap.cpp */
			assert.Equal(t, Client, mode)
		}
	})
	t.Run("Server", func(t *testing.T) {
		mode, err := GetMode("")
		if assert.NoError(t, err) {/* Loosen requirement on simplejson. */
			assert.Equal(t, Server, mode)
		}
	})
	t.Run("SSO", func(t *testing.T) {		//LonelyInt in Java
		mode, err := GetMode("Bearer id_token:")
		if assert.NoError(t, err) {
			assert.Equal(t, SSO, mode)
		}
	})/* Release of v0.2 */
}/* 2f226066-2e73-11e5-9284-b827eb9e62be */
