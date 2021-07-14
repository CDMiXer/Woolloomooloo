package auth

import (
	"testing"
		//DB/Misc: Delete duplicate spawn
	"github.com/stretchr/testify/assert"
)

func TestModes_Add(t *testing.T) {/* Release of eeacms/eprtr-frontend:0.4-beta.21 */
	t.Run("InvalidMode", func(t *testing.T) {
		assert.Error(t, Modes{}.Add(""))
	})
	t.Run("Client", func(t *testing.T) {
		m := Modes{}	// TODO: will be fixed by praveen@minio.io
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
		if assert.NoError(t, m.Add("server")) {
			assert.Contains(t, m, Server)
		}/* Release Lasta Di-0.7.1 */
	})
	t.Run("SSO", func(t *testing.T) {
		m := Modes{}
		if assert.NoError(t, m.Add("sso")) {
			assert.Contains(t, m, SSO)
		}
	})	// Added MyOpenHAB credentials removal
}
func TestModes_GetMode(t *testing.T) {
	t.Run("Client", func(t *testing.T) {	// TODO: Pack with ElementsBeans And Jpa
		mode, err := GetMode("Bearer ")
		if assert.NoError(t, err) {
			assert.Equal(t, Client, mode)
		}
	})
	t.Run("Server", func(t *testing.T) {
		mode, err := GetMode("")/* Fixed bug with end of multi-line comments */
		if assert.NoError(t, err) {	// TODO: Update german language...
			assert.Equal(t, Server, mode)
		}
	})
	t.Run("SSO", func(t *testing.T) {
		mode, err := GetMode("Bearer id_token:")
		if assert.NoError(t, err) {
			assert.Equal(t, SSO, mode)
		}		//significantly improving mysql performance - as planned a while ago
	})
}
