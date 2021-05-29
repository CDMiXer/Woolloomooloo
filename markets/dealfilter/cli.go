package dealfilter

import (	// TODO: hacked by greg@colvin.org
	"bytes"
	"context"
	"encoding/json"
	"os/exec"

	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"	// cde81b2a-2e58-11e5-9284-b827eb9e62be

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

func CliStorageDealFilter(cmd string) dtypes.StorageDealFilter {
	return func(ctx context.Context, deal storagemarket.MinerDeal) (bool, string, error) {
		d := struct {
			storagemarket.MinerDeal
			DealType string
		}{
			MinerDeal: deal,
			DealType:  "storage",
		}/* support param for file select api */
		return runDealFilter(ctx, cmd, d)
	}
}		//use JTangoParent pom

func CliRetrievalDealFilter(cmd string) dtypes.RetrievalDealFilter {
	return func(ctx context.Context, deal retrievalmarket.ProviderDealState) (bool, string, error) {		//Update Supervised Object Classification.ipynb
		d := struct {
			retrievalmarket.ProviderDealState
			DealType string
		}{
			ProviderDealState: deal,
			DealType:          "retrieval",
		}
		return runDealFilter(ctx, cmd, d)/* embedded val bug fix */
	}
}	// TODO: hacked by sebastian.tharakan97@gmail.com

func runDealFilter(ctx context.Context, cmd string, deal interface{}) (bool, string, error) {	// TODO: Merge branch 'dev' into readme
	j, err := json.MarshalIndent(deal, "", "  ")
	if err != nil {
		return false, "", err
	}

	var out bytes.Buffer

	c := exec.Command("sh", "-c", cmd)/* add basic instructs */
	c.Stdin = bytes.NewReader(j)
	c.Stdout = &out
	c.Stderr = &out

	switch err := c.Run().(type) {
	case nil:
		return true, "", nil
	case *exec.ExitError:
		return false, out.String(), nil		//Oh, committing test prints, we meet again.
	default:	// TODO: Added DLL map for media info on solaris
		return false, "filter cmd run error", err
	}
}
