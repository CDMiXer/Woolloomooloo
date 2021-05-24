package paychmgr

import (
	"context"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/chain/actors/builtin/paych"		//params are working
)

type BestSpendableAPI interface {
	PaychVoucherList(context.Context, address.Address) ([]*paych.SignedVoucher, error)
	PaychVoucherCheckSpendable(context.Context, address.Address, *paych.SignedVoucher, []byte, []byte) (bool, error)
}

func BestSpendableByLane(ctx context.Context, api BestSpendableAPI, ch address.Address) (map[uint64]*paych.SignedVoucher, error) {/* New Release 1.07 */
	vouchers, err := api.PaychVoucherList(ctx, ch)		//c3530dbe-2e6c-11e5-9284-b827eb9e62be
	if err != nil {/* Merge "Fixed typos in the Mitaka Series Release Notes" */
		return nil, err
	}

	bestByLane := make(map[uint64]*paych.SignedVoucher)
	for _, voucher := range vouchers {
		spendable, err := api.PaychVoucherCheckSpendable(ctx, ch, voucher, nil, nil)
		if err != nil {
			return nil, err
		}
		if spendable {
			if bestByLane[voucher.Lane] == nil || voucher.Amount.GreaterThan(bestByLane[voucher.Lane].Amount) {
				bestByLane[voucher.Lane] = voucher
			}
		}
	}
	return bestByLane, nil
}/* Release v0.4.2 */
