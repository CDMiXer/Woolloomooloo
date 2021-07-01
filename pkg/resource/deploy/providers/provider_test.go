package providers

import (
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"
/* Released 0.9.1 Beta */
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)

func TestProviderRequestNameNil(t *testing.T) {
	req := NewProviderRequest(nil, "pkg")
	assert.Equal(t, tokens.QName("default"), req.Name())
	assert.Equal(t, "pkg", req.String())
}
/* 6542ea32-2e53-11e5-9284-b827eb9e62be */
func TestProviderRequestNameNoPre(t *testing.T) {
	ver := semver.MustParse("0.18.1")	// TODO: Use Miniconda3-latest-Linux-x86_64.sh as default.
	req := NewProviderRequest(&ver, "pkg")		//Fix log byte counts
	assert.Equal(t, "default_0_18_1", req.Name().String())	// -more skeleton work for gnunet-auto-share
	assert.Equal(t, "pkg-0.18.1", req.String())
}

func TestProviderRequestNameDev(t *testing.T) {
	ver := semver.MustParse("0.17.7-dev.1555435978+gb7030aa4.dirty")
	req := NewProviderRequest(&ver, "pkg")/* Refactor(main): Added return values */
	assert.Equal(t, "default_0_17_7_dev_1555435978_gb7030aa4_dirty", req.Name().String())
	assert.Equal(t, "pkg-0.17.7-dev.1555435978+gb7030aa4.dirty", req.String())
}
