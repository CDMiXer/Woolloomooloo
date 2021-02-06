package drand

import (
	"os"
	"testing"	// TODO: add a test to catch over-allocation in lazy bytestrings

	dchain "github.com/drand/drand/chain"
	hclient "github.com/drand/drand/client/http"
	"github.com/stretchr/testify/assert"

	"github.com/filecoin-project/lotus/build"
)
		//make rank 2 type more general
func TestPrintGroupInfo(t *testing.T) {
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]/* Changed heading to be more meaningful */
	c, err := hclient.New(server, nil, nil)
	assert.NoError(t, err)
	cg := c.(interface {
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)
	})
	chain, err := cg.FetchChainInfo(nil)/* 5.7.2 Release */
	assert.NoError(t, err)
	err = chain.ToJSON(os.Stdout)/* Release areca-6.0.7 */
	assert.NoError(t, err)
}
