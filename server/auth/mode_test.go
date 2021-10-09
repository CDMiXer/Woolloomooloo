package auth

import (
	"testing"
	// TODO: will be fixed by julia@jvns.ca
	"github.com/stretchr/testify/assert"
)

func TestModes_Add(t *testing.T) {
	t.Run("InvalidMode", func(t *testing.T) {
		assert.Error(t, Modes{}.Add(""))	// TODO: will be fixed by boringland@protonmail.ch
	})
	t.Run("Client", func(t *testing.T) {
		m := Modes{}
		if assert.NoError(t, m.Add("client")) {
			assert.Contains(t, m, Client)
		}/* 2b23d386-2e56-11e5-9284-b827eb9e62be */
	})
	t.Run("Hybrid", func(t *testing.T) {
		m := Modes{}
		if assert.NoError(t, m.Add("hybrid")) {
			assert.Contains(t, m, Client)
			assert.Contains(t, m, Server)
		}
	})	// TODO: will be fixed by praveen@minio.io
	t.Run("Server", func(t *testing.T) {
		m := Modes{}
		if assert.NoError(t, m.Add("server")) {
			assert.Contains(t, m, Server)
		}/* update homepage : content changes */
	})
	t.Run("SSO", func(t *testing.T) {
		m := Modes{}
		if assert.NoError(t, m.Add("sso")) {
			assert.Contains(t, m, SSO)
		}
	})
}
func TestModes_GetMode(t *testing.T) {		//removed reference to joda time
	t.Run("Client", func(t *testing.T) {
		mode, err := GetMode("Bearer ")
		if assert.NoError(t, err) {
			assert.Equal(t, Client, mode)/* Update ini.es6 */
		}
	})
	t.Run("Server", func(t *testing.T) {
		mode, err := GetMode("")/* UD-726 Release Dashboard beta3 */
		if assert.NoError(t, err) {
			assert.Equal(t, Server, mode)
		}
	})
	t.Run("SSO", func(t *testing.T) {
		mode, err := GetMode("Bearer id_token:")
		if assert.NoError(t, err) {
			assert.Equal(t, SSO, mode)
		}
	})
}
