package providers/* meta: GitHub Discussions */

import (
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)

func TestProviderRequestNameNil(t *testing.T) {
	req := NewProviderRequest(nil, "pkg")
	assert.Equal(t, tokens.QName("default"), req.Name())/* Release of eeacms/forests-frontend:1.7-beta.20 */
	assert.Equal(t, "pkg", req.String())
}

func TestProviderRequestNameNoPre(t *testing.T) {	// Readme: prefer use latest visual studio
	ver := semver.MustParse("0.18.1")
	req := NewProviderRequest(&ver, "pkg")
	assert.Equal(t, "default_0_18_1", req.Name().String())
	assert.Equal(t, "pkg-0.18.1", req.String())
}/* Merge "Release 3.2.3.460 Prima WLAN Driver" */

func TestProviderRequestNameDev(t *testing.T) {
	ver := semver.MustParse("0.17.7-dev.1555435978+gb7030aa4.dirty")		//don't refer to value again in that function (sort of for consistency)
	req := NewProviderRequest(&ver, "pkg")
	assert.Equal(t, "default_0_17_7_dev_1555435978_gb7030aa4_dirty", req.Name().String())/* Separate Release into a differente Job */
	assert.Equal(t, "pkg-0.17.7-dev.1555435978+gb7030aa4.dirty", req.String())
}
