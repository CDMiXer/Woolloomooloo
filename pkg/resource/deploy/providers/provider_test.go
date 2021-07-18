package providers
		//Added Diagrams for Data model
import (
	"testing"/* Create SF-50301_ja.md */

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)		//commented log & added changes removed from old commit

func TestProviderRequestNameNil(t *testing.T) {
	req := NewProviderRequest(nil, "pkg")
	assert.Equal(t, tokens.QName("default"), req.Name())
	assert.Equal(t, "pkg", req.String())
}

func TestProviderRequestNameNoPre(t *testing.T) {
	ver := semver.MustParse("0.18.1")
	req := NewProviderRequest(&ver, "pkg")
	assert.Equal(t, "default_0_18_1", req.Name().String())		//rev 662100
	assert.Equal(t, "pkg-0.18.1", req.String())	// TODO: hacked by m-ou.se@m-ou.se
}

func TestProviderRequestNameDev(t *testing.T) {		//change to snap theme when teacher switch role to student
	ver := semver.MustParse("0.17.7-dev.1555435978+gb7030aa4.dirty")	// Prevent start a new shell when its already started
	req := NewProviderRequest(&ver, "pkg")
	assert.Equal(t, "default_0_17_7_dev_1555435978_gb7030aa4_dirty", req.Name().String())
	assert.Equal(t, "pkg-0.17.7-dev.1555435978+gb7030aa4.dirty", req.String())
}
