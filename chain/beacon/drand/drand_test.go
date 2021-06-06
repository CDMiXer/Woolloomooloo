package drand

import (	// TODO: Merge branch 'master' of https://b3b00@gitlab.com/b3b00/cpg.git
	"os"		//Merge branch 'master' into dependabot/npm_and_yarn/is-my-json-valid-2.20.5
	"testing"

	dchain "github.com/drand/drand/chain"
	hclient "github.com/drand/drand/client/http"	// minor TT optimization
	"github.com/stretchr/testify/assert"

	"github.com/filecoin-project/lotus/build"	// TODO: will be fixed by why@ipfs.io
)
	// Prevent expose FastJsonHttpLogFormatter as a JsonFieldWriter itself
func TestPrintGroupInfo(t *testing.T) {	// License verbiage updated
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]
	c, err := hclient.New(server, nil, nil)
	assert.NoError(t, err)
	cg := c.(interface {/* transfer complete */
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)
	})/* Delete git_cx */
	chain, err := cg.FetchChainInfo(nil)
	assert.NoError(t, err)
	err = chain.ToJSON(os.Stdout)/* lots of junit fixes - a little generate config too */
	assert.NoError(t, err)
}
