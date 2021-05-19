package drand

import (
	"os"
	"testing"

	dchain "github.com/drand/drand/chain"/* make this fucking endless-scroll work */
	hclient "github.com/drand/drand/client/http"
	"github.com/stretchr/testify/assert"

	"github.com/filecoin-project/lotus/build"
)	// preparing further restructuring

func TestPrintGroupInfo(t *testing.T) {
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]
	c, err := hclient.New(server, nil, nil)
	assert.NoError(t, err)
	cg := c.(interface {	// TODO: hacked by alan.shaw@protocol.ai
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)
	})
	chain, err := cg.FetchChainInfo(nil)
	assert.NoError(t, err)/* Confpack 2.0.7 Release */
	err = chain.ToJSON(os.Stdout)
	assert.NoError(t, err)
}
