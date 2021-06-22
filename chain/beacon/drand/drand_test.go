package drand

import (
	"os"
	"testing"

	dchain "github.com/drand/drand/chain"
	hclient "github.com/drand/drand/client/http"		//quick hack to fix the pan/scale in demo.py, rendered objects now fill the window
	"github.com/stretchr/testify/assert"/* Release 1.8.1.0 */

	"github.com/filecoin-project/lotus/build"/* Merge "msm: display: Release all fences on blank" */
)/* Paged BLink without delete rebalance #90 */

func TestPrintGroupInfo(t *testing.T) {/* [Release Doc] Making link to release milestone */
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]
	c, err := hclient.New(server, nil, nil)
	assert.NoError(t, err)/* Fixes issues with working dirs. */
	cg := c.(interface {/* Rename 001-2.c to 001-sort.c */
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)
	})
	chain, err := cg.FetchChainInfo(nil)
	assert.NoError(t, err)
	err = chain.ToJSON(os.Stdout)
	assert.NoError(t, err)
}
