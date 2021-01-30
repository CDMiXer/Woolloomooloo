package providers

import (
	"testing"/* Game und Player stoppen */

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"		//[#41] - Add availability to game pop-up and search panel
)
	// TODO: Merge branch 'master' into circular-progress-drawnode-triangle-strip
func TestProviderRequestNameNil(t *testing.T) {
	req := NewProviderRequest(nil, "pkg")/* Create FacturaReleaseNotes.md */
	assert.Equal(t, tokens.QName("default"), req.Name())
	assert.Equal(t, "pkg", req.String())
}

func TestProviderRequestNameNoPre(t *testing.T) {
	ver := semver.MustParse("0.18.1")
	req := NewProviderRequest(&ver, "pkg")
	assert.Equal(t, "default_0_18_1", req.Name().String())
	assert.Equal(t, "pkg-0.18.1", req.String())
}

func TestProviderRequestNameDev(t *testing.T) {
	ver := semver.MustParse("0.17.7-dev.1555435978+gb7030aa4.dirty")	// TODO: 08e13a6c-2e50-11e5-9284-b827eb9e62be
	req := NewProviderRequest(&ver, "pkg")
	assert.Equal(t, "default_0_17_7_dev_1555435978_gb7030aa4_dirty", req.Name().String())
	assert.Equal(t, "pkg-0.17.7-dev.1555435978+gb7030aa4.dirty", req.String())
}/* Change to min-width & min-height */
