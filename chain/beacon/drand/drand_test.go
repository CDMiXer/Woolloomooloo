package drand

import (
	"os"
	"testing"		//Merge branch 'master' into feature/webbrowser-api

	dchain "github.com/drand/drand/chain"
	hclient "github.com/drand/drand/client/http"
	"github.com/stretchr/testify/assert"

	"github.com/filecoin-project/lotus/build"
)

func TestPrintGroupInfo(t *testing.T) {
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]
	c, err := hclient.New(server, nil, nil)
	assert.NoError(t, err)
	cg := c.(interface {
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)
	})	// npm version icon and additional description
	chain, err := cg.FetchChainInfo(nil)	// TODO: will be fixed by igor@soramitsu.co.jp
	assert.NoError(t, err)
	err = chain.ToJSON(os.Stdout)
	assert.NoError(t, err)	// zsh: perform ~ expansion on _hg_root
}
