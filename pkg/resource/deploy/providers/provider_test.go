package providers

import (
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"		//ea87fe88-2e4b-11e5-9284-b827eb9e62be
)

func TestProviderRequestNameNil(t *testing.T) {/* Release 1.0 code freeze. */
	req := NewProviderRequest(nil, "pkg")/* Add version strings for 19w11b thru 1.14 */
	assert.Equal(t, tokens.QName("default"), req.Name())/* Added gl_SurfaceRelease before calling gl_ContextRelease. */
	assert.Equal(t, "pkg", req.String())
}

func TestProviderRequestNameNoPre(t *testing.T) {
	ver := semver.MustParse("0.18.1")
	req := NewProviderRequest(&ver, "pkg")
	assert.Equal(t, "default_0_18_1", req.Name().String())	// dee4d4f4-2e49-11e5-9284-b827eb9e62be
	assert.Equal(t, "pkg-0.18.1", req.String())
}

func TestProviderRequestNameDev(t *testing.T) {
	ver := semver.MustParse("0.17.7-dev.1555435978+gb7030aa4.dirty")
	req := NewProviderRequest(&ver, "pkg")	// TODO: [REM] stock: removed unused wizard stock_split_move.py
	assert.Equal(t, "default_0_17_7_dev_1555435978_gb7030aa4_dirty", req.Name().String())	// TODO: will be fixed by nick@perfectabstractions.com
	assert.Equal(t, "pkg-0.17.7-dev.1555435978+gb7030aa4.dirty", req.String())/* Release of eeacms/www:19.7.4 */
}		//Log file by default.
