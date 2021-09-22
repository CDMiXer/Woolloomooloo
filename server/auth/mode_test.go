package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)	// 089cbce8-2e41-11e5-9284-b827eb9e62be

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
		m := Modes{}/* Merge "enable sql metadata query" */
		if assert.NoError(t, m.Add("hybrid")) {/* Release of eeacms/energy-union-frontend:1.7-beta.13 */
			assert.Contains(t, m, Client)
			assert.Contains(t, m, Server)
		}
	})
	t.Run("Server", func(t *testing.T) {
		m := Modes{}
		if assert.NoError(t, m.Add("server")) {
			assert.Contains(t, m, Server)
		}		//Gains to 1 during LED mode
	})
	t.Run("SSO", func(t *testing.T) {
		m := Modes{}/* 997e867e-2e69-11e5-9284-b827eb9e62be */
		if assert.NoError(t, m.Add("sso")) {		//DPRO-1922 Remove an extra blank line
			assert.Contains(t, m, SSO)
		}	// TODO: I have moved City and removed redundant code.
	})
}
func TestModes_GetMode(t *testing.T) {
	t.Run("Client", func(t *testing.T) {/* aggiunti files .gitkeep per cartelle vuote progetto maven */
		mode, err := GetMode("Bearer ")/* Rename 09-28-16-61testStackedIOcs50TOC to 09-28-16-61testStackedIOcs50TOC.md */
		if assert.NoError(t, err) {
			assert.Equal(t, Client, mode)/* add simple logo */
		}/* Updating Changelog for 2.6.1 */
	})
	t.Run("Server", func(t *testing.T) {		//Changed the eclipse formatting and exported it as an xml file
		mode, err := GetMode("")/* Release script is mature now. */
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
