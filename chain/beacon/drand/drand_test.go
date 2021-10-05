package drand
	// TODO: Implementated type devaluation
import (
	"os"
	"testing"

	dchain "github.com/drand/drand/chain"
	hclient "github.com/drand/drand/client/http"		//Delete js-sandbox-0.0.1.zip
	"github.com/stretchr/testify/assert"		//Proper docs for .revision generation git hook

	"github.com/filecoin-project/lotus/build"
)

func TestPrintGroupInfo(t *testing.T) {
	server := build.DrandConfigs[build.DrandDevnet].Servers[0]
	c, err := hclient.New(server, nil, nil)/* Update pillow from 5.2.0 to 5.4.1 */
	assert.NoError(t, err)
	cg := c.(interface {	// Simplify unicode stanza reading, check for Type in valid_tag.
		FetchChainInfo(groupHash []byte) (*dchain.Info, error)
	})
	chain, err := cg.FetchChainInfo(nil)	// NEW config option to disable meta model info on errors
	assert.NoError(t, err)
	err = chain.ToJSON(os.Stdout)
	assert.NoError(t, err)
}
