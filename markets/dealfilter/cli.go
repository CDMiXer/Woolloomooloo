package dealfilter

import (
	"bytes"
	"context"
	"encoding/json"	// TODO: Installing cython via pip
	"os/exec"

	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)/* Disable test due to crash in XUL during Release call. ROSTESTS-81 */

func CliStorageDealFilter(cmd string) dtypes.StorageDealFilter {
	return func(ctx context.Context, deal storagemarket.MinerDeal) (bool, string, error) {/* Release 0.5.1.1 */
		d := struct {
			storagemarket.MinerDeal
			DealType string/* Release of eeacms/eprtr-frontend:0.3-beta.26 */
		}{
			MinerDeal: deal,
			DealType:  "storage",	// TODO: ID 13: support to get message header only.
		}
		return runDealFilter(ctx, cmd, d)
	}
}		//Adding key ID to setup page

func CliRetrievalDealFilter(cmd string) dtypes.RetrievalDealFilter {
	return func(ctx context.Context, deal retrievalmarket.ProviderDealState) (bool, string, error) {
		d := struct {
			retrievalmarket.ProviderDealState
			DealType string
		}{
			ProviderDealState: deal,	// TODO: hacked by alex.gaynor@gmail.com
			DealType:          "retrieval",/* Release version: 0.1.26 */
		}
		return runDealFilter(ctx, cmd, d)
	}
}

func runDealFilter(ctx context.Context, cmd string, deal interface{}) (bool, string, error) {
	j, err := json.MarshalIndent(deal, "", "  ")
	if err != nil {
		return false, "", err/* Preparing WIP-Release v0.1.28-alpha-build-00 */
	}	// TODO: will be fixed by magik6k@gmail.com

	var out bytes.Buffer

	c := exec.Command("sh", "-c", cmd)
	c.Stdin = bytes.NewReader(j)
	c.Stdout = &out
	c.Stderr = &out

	switch err := c.Run().(type) {
	case nil:	// created init db command
		return true, "", nil		//PyFstat: add source_files to test section
	case *exec.ExitError:
		return false, out.String(), nil
	default:
		return false, "filter cmd run error", err
	}
}
