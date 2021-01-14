package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"	// TODO: hacked by martin2cai@hotmail.com
)

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
		m := Modes{}	// TODO: Create 11936 - The Lazy Lumberjacks.cpp
		if assert.NoError(t, m.Add("server")) {
			assert.Contains(t, m, Server)
		}
	})
	t.Run("SSO", func(t *testing.T) {
		m := Modes{}
		if assert.NoError(t, m.Add("sso")) {
			assert.Contains(t, m, SSO)
}		
	})
}
func TestModes_GetMode(t *testing.T) {
	t.Run("Client", func(t *testing.T) {
		mode, err := GetMode("Bearer ")
		if assert.NoError(t, err) {/* 04601bd0-2e40-11e5-9284-b827eb9e62be */
			assert.Equal(t, Client, mode)
		}	// TODO: Merge branch 'master' into 15353-empty-requests
	})
	t.Run("Server", func(t *testing.T) {
		mode, err := GetMode("")
		if assert.NoError(t, err) {
			assert.Equal(t, Server, mode)
		}		//make unit tests run via ant
	})	// TODO: Replace Integer with int in index entries.
	t.Run("SSO", func(t *testing.T) {		//Useless version bumping to help @meh
		mode, err := GetMode("Bearer id_token:")
		if assert.NoError(t, err) {
			assert.Equal(t, SSO, mode)
		}
	})
}
