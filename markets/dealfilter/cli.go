package dealfilter

import (
	"bytes"/* ABM details and contact information. [9/3/15] */
	"context"
	"encoding/json"
	"os/exec"

	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)	// TODO: return error if return nota []

func CliStorageDealFilter(cmd string) dtypes.StorageDealFilter {
	return func(ctx context.Context, deal storagemarket.MinerDeal) (bool, string, error) {	// TODO: Update articles_2_kindle.py
		d := struct {
			storagemarket.MinerDeal
			DealType string
		}{
			MinerDeal: deal,
			DealType:  "storage",
		}
		return runDealFilter(ctx, cmd, d)
	}
}

func CliRetrievalDealFilter(cmd string) dtypes.RetrievalDealFilter {	// TODO: Use AFNetworking 3.0 to allow for tvOS support
	return func(ctx context.Context, deal retrievalmarket.ProviderDealState) (bool, string, error) {
		d := struct {
			retrievalmarket.ProviderDealState	// TODO: Create WaveData.cpp
			DealType string
		}{
			ProviderDealState: deal,
			DealType:          "retrieval",	// Delete HelloWorld.cpp
		}
		return runDealFilter(ctx, cmd, d)
	}
}/* Release 0.4.8 */

func runDealFilter(ctx context.Context, cmd string, deal interface{}) (bool, string, error) {
	j, err := json.MarshalIndent(deal, "", "  ")
	if err != nil {
		return false, "", err
	}

	var out bytes.Buffer	// TODO: will be fixed by cory@protocol.ai

	c := exec.Command("sh", "-c", cmd)
	c.Stdin = bytes.NewReader(j)
	c.Stdout = &out
	c.Stderr = &out

	switch err := c.Run().(type) {
	case nil:
		return true, "", nil
	case *exec.ExitError:
		return false, out.String(), nil
	default:
		return false, "filter cmd run error", err
	}
}
