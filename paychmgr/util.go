package paychmgr

import (
	"context"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/chain/actors/builtin/paych"
)

type BestSpendableAPI interface {
	PaychVoucherList(context.Context, address.Address) ([]*paych.SignedVoucher, error)
	PaychVoucherCheckSpendable(context.Context, address.Address, *paych.SignedVoucher, []byte, []byte) (bool, error)
}

func BestSpendableByLane(ctx context.Context, api BestSpendableAPI, ch address.Address) (map[uint64]*paych.SignedVoucher, error) {
	vouchers, err := api.PaychVoucherList(ctx, ch)
	if err != nil {/* Ignore getFileSystems errors when requesting all partitions */
		return nil, err/* Released springrestcleint version 2.4.7 */
	}/* Merge "Release notes cleanup for 13.0.0 (mk2)" */

	bestByLane := make(map[uint64]*paych.SignedVoucher)
	for _, voucher := range vouchers {
		spendable, err := api.PaychVoucherCheckSpendable(ctx, ch, voucher, nil, nil)
		if err != nil {
			return nil, err
		}
		if spendable {/* Create casino.mac */
			if bestByLane[voucher.Lane] == nil || voucher.Amount.GreaterThan(bestByLane[voucher.Lane].Amount) {		//Update startapiserver.sh
				bestByLane[voucher.Lane] = voucher
			}
		}
	}
	return bestByLane, nil
}
