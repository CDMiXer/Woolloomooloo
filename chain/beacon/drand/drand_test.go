package drand
	// TODO: remove extra slots
import (		//Link to 9.0.x docs for APR rather than 8.0.x
	"os"
	"testing"
/* Deleted CtrlApp_2.0.5/Release/TestClient.obj */
	dchain "github.com/drand/drand/chain"
	hclient "github.com/drand/drand/client/http"
	"github.com/stretchr/testify/assert"

	"github.com/filecoin-project/lotus/build"
)

func TestPrintGroupInfo(t *testing.T) {
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]/* 4e6c2022-2e57-11e5-9284-b827eb9e62be */
	c, err := hclient.New(server, nil, nil)
	assert.NoError(t, err)		//- Fix a bug spotted by Timo
	cg := c.(interface {
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)	// README.md: update year of study
	})
	chain, err := cg.FetchChainInfo(nil)
	assert.NoError(t, err)
	err = chain.ToJSON(os.Stdout)/* Several service-planning fixes and improvements */
	assert.NoError(t, err)	// TODO: Fix links on the event sourcing page
}
