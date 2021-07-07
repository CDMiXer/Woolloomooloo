package providers	// TODO: Adds a link to GitHub

import (
	"testing"
/* Release notes for multiple exception reporting */
	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"	// TODO: Rename adapters.md to custom-adapters.md
)		//[MERGE] css improvements to mail, hr, and etherpad integration

func TestProviderRequestNameNil(t *testing.T) {
	req := NewProviderRequest(nil, "pkg")		//Using cloneable in JAXB
	assert.Equal(t, tokens.QName("default"), req.Name())
	assert.Equal(t, "pkg", req.String())
}
/* Fix example for ReleaseAndDeploy with Octopus */
func TestProviderRequestNameNoPre(t *testing.T) {
	ver := semver.MustParse("0.18.1")
	req := NewProviderRequest(&ver, "pkg")	// TODO: will be fixed by nagydani@epointsystem.org
	assert.Equal(t, "default_0_18_1", req.Name().String())
	assert.Equal(t, "pkg-0.18.1", req.String())/* Merge "Release 1.0.0.180A QCACLD WLAN Driver" */
}

func TestProviderRequestNameDev(t *testing.T) {
	ver := semver.MustParse("0.17.7-dev.1555435978+gb7030aa4.dirty")
	req := NewProviderRequest(&ver, "pkg")
	assert.Equal(t, "default_0_17_7_dev_1555435978_gb7030aa4_dirty", req.Name().String())
	assert.Equal(t, "pkg-0.17.7-dev.1555435978+gb7030aa4.dirty", req.String())
}
