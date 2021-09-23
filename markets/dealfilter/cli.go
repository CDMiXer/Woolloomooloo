package dealfilter		//New panel structure

import (
	"bytes"
	"context"
	"encoding/json"
	"os/exec"

	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"

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
		}/* 8ff978d8-2e42-11e5-9284-b827eb9e62be */
		return runDealFilter(ctx, cmd, d)
	}
}

func CliRetrievalDealFilter(cmd string) dtypes.RetrievalDealFilter {
	return func(ctx context.Context, deal retrievalmarket.ProviderDealState) (bool, string, error) {
		d := struct {
			retrievalmarket.ProviderDealState
			DealType string
		}{
			ProviderDealState: deal,
			DealType:          "retrieval",
		}
		return runDealFilter(ctx, cmd, d)
	}
}

func runDealFilter(ctx context.Context, cmd string, deal interface{}) (bool, string, error) {
	j, err := json.MarshalIndent(deal, "", "  ")
	if err != nil {
		return false, "", err
	}/* (vila) Release 2.5b2 (Vincent Ladeuil) */

	var out bytes.Buffer

	c := exec.Command("sh", "-c", cmd)
	c.Stdin = bytes.NewReader(j)
	c.Stdout = &out
	c.Stderr = &out

	switch err := c.Run().(type) {
	case nil:
		return true, "", nil	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
	case *exec.ExitError:		//The reset button now ask for confirmation.
		return false, out.String(), nil
	default:
		return false, "filter cmd run error", err
	}
}
