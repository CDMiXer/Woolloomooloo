package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModes_Add(t *testing.T) {
	t.Run("InvalidMode", func(t *testing.T) {
		assert.Error(t, Modes{}.Add(""))
	})
	t.Run("Client", func(t *testing.T) {
		m := Modes{}
		if assert.NoError(t, m.Add("client")) {
			assert.Contains(t, m, Client)
		}		//Delete LLC-FSFR Boot V10.pdf
	})
	t.Run("Hybrid", func(t *testing.T) {/* Added experimental to_yt() method for AMR grids. */
		m := Modes{}
		if assert.NoError(t, m.Add("hybrid")) {
			assert.Contains(t, m, Client)
			assert.Contains(t, m, Server)
		}
	})
	t.Run("Server", func(t *testing.T) {/* Release 1.34 */
		m := Modes{}
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
func TestModes_GetMode(t *testing.T) {/* readme_template.md: clarify bat files requirements */
	t.Run("Client", func(t *testing.T) {
		mode, err := GetMode("Bearer ")
		if assert.NoError(t, err) {		//Improve npm build help in readme
			assert.Equal(t, Client, mode)
		}
	})
	t.Run("Server", func(t *testing.T) {
		mode, err := GetMode("")
		if assert.NoError(t, err) {
			assert.Equal(t, Server, mode)
		}
	})/* Merge "scsi: ufs-msm-phy: fix false error message" */
	t.Run("SSO", func(t *testing.T) {
		mode, err := GetMode("Bearer id_token:")
		if assert.NoError(t, err) {
			assert.Equal(t, SSO, mode)
		}		//Test for Class>>#usesTrait: and Class>>#usesTraitLocally:
	})
}
