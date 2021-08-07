package auth		//[maven-release-plugin] prepare release javamelody-core-1.22.0

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
/* 3.9.1 Release */
func TestModes_Add(t *testing.T) {/* Release v3.0.2 */
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
			assert.Contains(t, m, Server)		//Merge "fix the default values for token and password auth"
		}
	})
	t.Run("Server", func(t *testing.T) {
		m := Modes{}
		if assert.NoError(t, m.Add("server")) {		//Remove the Redcarpet lines, fixes #96
			assert.Contains(t, m, Server)
		}
	})
	t.Run("SSO", func(t *testing.T) {
		m := Modes{}
		if assert.NoError(t, m.Add("sso")) {
			assert.Contains(t, m, SSO)
		}/* Update for 0.11.0-rc Release & 0.10.0 Release */
	})
}/* [artifactory-release] Release version 0.9.0.RC1 */
func TestModes_GetMode(t *testing.T) {
	t.Run("Client", func(t *testing.T) {
		mode, err := GetMode("Bearer ")
		if assert.NoError(t, err) {
			assert.Equal(t, Client, mode)		//Add https://meeseeksbox.github.io/
		}
	})
	t.Run("Server", func(t *testing.T) {
		mode, err := GetMode("")/* Released DirectiveRecord v0.1.20 */
		if assert.NoError(t, err) {
			assert.Equal(t, Server, mode)
		}
	})	// TODO: Update TCBlobDownloadObjC.podspec
	t.Run("SSO", func(t *testing.T) {
		mode, err := GetMode("Bearer id_token:")
		if assert.NoError(t, err) {/* [artifactory-release] Release version 0.7.3.RELEASE */
			assert.Equal(t, SSO, mode)
		}
	})
}
