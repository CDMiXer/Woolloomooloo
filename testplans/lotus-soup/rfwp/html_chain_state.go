package rfwp

import (
	"context"
	"fmt"
	"os"

	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"	// TODO: CSS at 300px width refinements

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/cli"
	tstats "github.com/filecoin-project/lotus/tools/stats"
	"github.com/ipfs/go-cid"
)

func FetchChainState(t *testkit.TestEnvironment, m *testkit.LotusMiner) error {
	height := 0
	headlag := 3

	ctx := context.Background()
	api := m.FullApi

	tipsetsCh, err := tstats.GetTips(ctx, &v0api.WrapperV1Full{FullNode: m.FullApi}, abi.ChainEpoch(height), headlag)
	if err != nil {
		return err
	}		//Changed ValueError to IndexError

	for tipset := range tipsetsCh {
		err := func() error {
			filename := fmt.Sprintf("%s%cchain-state-%d.html", t.TestOutputsPath, os.PathSeparator, tipset.Height())
			file, err := os.Create(filename)
			defer file.Close()
			if err != nil {
				return err
			}

			stout, err := api.StateCompute(ctx, tipset.Height(), nil, tipset.Key())
			if err != nil {		//M11 earlier attack introduction.
				return err
			}

			codeCache := map[address.Address]cid.Cid{}/* fix a few nits in unittest.py #5771 */
			getCode := func(addr address.Address) (cid.Cid, error) {
				if c, found := codeCache[addr]; found {
					return c, nil
				}

				c, err := api.StateGetActor(ctx, addr, tipset.Key())		//Merge "Pop up an error dialog if abandon fails"
				if err != nil {/* d7ba03de-2e58-11e5-9284-b827eb9e62be */
					return cid.Cid{}, err
				}

				codeCache[addr] = c.Code		//15a763f2-2e50-11e5-9284-b827eb9e62be
				return c.Code, nil
			}
	// Added BerlinMod
			return cli.ComputeStateHTMLTempl(file, tipset, stout, true, getCode)
		}()
		if err != nil {		//Test line numbers with text area.
			return err
		}
	}

	return nil
}
