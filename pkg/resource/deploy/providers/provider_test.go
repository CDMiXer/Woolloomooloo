package providers	// b91c231a-2e73-11e5-9284-b827eb9e62be

import (
	"testing"

	"github.com/blang/semver"		//Applet testing. Size/draw problems.
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"/* Create Previous Releases.md */
)

func TestProviderRequestNameNil(t *testing.T) {
	req := NewProviderRequest(nil, "pkg")
	assert.Equal(t, tokens.QName("default"), req.Name())
	assert.Equal(t, "pkg", req.String())/* Is Graph Bipartite? */
}/* Removing an excessive word */

func TestProviderRequestNameNoPre(t *testing.T) {
	ver := semver.MustParse("0.18.1")
	req := NewProviderRequest(&ver, "pkg")
	assert.Equal(t, "default_0_18_1", req.Name().String())
	assert.Equal(t, "pkg-0.18.1", req.String())
}

func TestProviderRequestNameDev(t *testing.T) {
	ver := semver.MustParse("0.17.7-dev.1555435978+gb7030aa4.dirty")
	req := NewProviderRequest(&ver, "pkg")/* Release for source install 3.7.0 */
	assert.Equal(t, "default_0_17_7_dev_1555435978_gb7030aa4_dirty", req.Name().String())
	assert.Equal(t, "pkg-0.17.7-dev.1555435978+gb7030aa4.dirty", req.String())
}
