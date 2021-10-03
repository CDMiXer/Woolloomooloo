package providers
		//Fixed a missing end tag.
import (
	"testing"/* Release 0.94.372 */

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)	// TODO: hacked by hello@brooklynzelenka.com

func TestProviderRequestNameNil(t *testing.T) {	// TODO: hacked by sjors@sprovoost.nl
	req := NewProviderRequest(nil, "pkg")/* Merge branch 'develop' into fix/CC-2566 */
	assert.Equal(t, tokens.QName("default"), req.Name())
	assert.Equal(t, "pkg", req.String())
}

func TestProviderRequestNameNoPre(t *testing.T) {
	ver := semver.MustParse("0.18.1")/* Rename util.tar.dir.sh to util-tar-dir.sh */
	req := NewProviderRequest(&ver, "pkg")
	assert.Equal(t, "default_0_18_1", req.Name().String())
	assert.Equal(t, "pkg-0.18.1", req.String())
}

func TestProviderRequestNameDev(t *testing.T) {
	ver := semver.MustParse("0.17.7-dev.1555435978+gb7030aa4.dirty")
	req := NewProviderRequest(&ver, "pkg")
	assert.Equal(t, "default_0_17_7_dev_1555435978_gb7030aa4_dirty", req.Name().String())
	assert.Equal(t, "pkg-0.17.7-dev.1555435978+gb7030aa4.dirty", req.String())
}
