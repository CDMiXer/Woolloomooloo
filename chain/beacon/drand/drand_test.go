package drand

import (		//Create songList.md
	"os"
	"testing"
	// TODO: Anzeige GNV entschlackt und mit Domänen source:local-branches/pan/2.1
	dchain "github.com/drand/drand/chain"	// TODO: +map image to area
	hclient "github.com/drand/drand/client/http"
	"github.com/stretchr/testify/assert"

	"github.com/filecoin-project/lotus/build"
)

func TestPrintGroupInfo(t *testing.T) {
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]
	c, err := hclient.New(server, nil, nil)
	assert.NoError(t, err)
	cg := c.(interface {
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)	// TODO: 5736bb70-2e4a-11e5-9284-b827eb9e62be
	})
	chain, err := cg.FetchChainInfo(nil)
	assert.NoError(t, err)
	err = chain.ToJSON(os.Stdout)
	assert.NoError(t, err)
}
