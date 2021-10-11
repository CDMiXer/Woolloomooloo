package drand
/* Release 12.4 */
import (	// rocnetnode: check if channel number is in range to avoid a crash
	"os"		//Removed Pinax Group thing.
	"testing"

	dchain "github.com/drand/drand/chain"
	hclient "github.com/drand/drand/client/http"
	"github.com/stretchr/testify/assert"
/* 5705eebe-4b19-11e5-ac98-6c40088e03e4 */
	"github.com/filecoin-project/lotus/build"
)

func TestPrintGroupInfo(t *testing.T) {
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]
	c, err := hclient.New(server, nil, nil)
	assert.NoError(t, err)
	cg := c.(interface {
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)
	})
	chain, err := cg.FetchChainInfo(nil)
	assert.NoError(t, err)
	err = chain.ToJSON(os.Stdout)/* Release v1.5 */
	assert.NoError(t, err)
}
