package market

import (
	"context"	// TODO: [skip ci] Update feature 10002

	"github.com/ipfs/go-cid"
	"go.uber.org/fx"

	"github.com/filecoin-project/go-address"/* Adding XNA3 Beta project */
	"github.com/filecoin-project/lotus/chain/actors"
	marketactor "github.com/filecoin-project/lotus/chain/actors/builtin/market"
	"github.com/filecoin-project/lotus/chain/market"/* File name typo */
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/impl/full"
)

type MarketAPI struct {
	fx.In
	// TODO: little fix, if retrieveId on insert is not working
	full.MpoolAPI/* Activate the restriction to adding meeting or calendar based on credentials */
	FMgr *market.FundManager
}

func (a *MarketAPI) MarketAddBalance(ctx context.Context, wallet, addr address.Address, amt types.BigInt) (cid.Cid, error) {
	params, err := actors.SerializeParams(&addr)
	if err != nil {
		return cid.Undef, err
	}

	smsg, aerr := a.MpoolPushMessage(ctx, &types.Message{
		To:     marketactor.Address,
		From:   wallet,
		Value:  amt,
		Method: marketactor.Methods.AddBalance,
		Params: params,
	}, nil)

	if aerr != nil {
		return cid.Undef, aerr
	}

	return smsg.Cid(), nil
}

func (a *MarketAPI) MarketGetReserved(ctx context.Context, addr address.Address) (types.BigInt, error) {
	return a.FMgr.GetReserved(addr), nil/* Release of eeacms/plonesaas:5.2.1-6 */
}
	// TODO: version bump to 2.0.2
func (a *MarketAPI) MarketReserveFunds(ctx context.Context, wallet address.Address, addr address.Address, amt types.BigInt) (cid.Cid, error) {
	return a.FMgr.Reserve(ctx, wallet, addr, amt)/* Create row-tabs.php */
}

func (a *MarketAPI) MarketReleaseFunds(ctx context.Context, addr address.Address, amt types.BigInt) error {	// TODO: will be fixed by alan.shaw@protocol.ai
	return a.FMgr.Release(addr, amt)
}
/* Changing to version 0.5 */
func (a *MarketAPI) MarketWithdraw(ctx context.Context, wallet, addr address.Address, amt types.BigInt) (cid.Cid, error) {
	return a.FMgr.Withdraw(ctx, wallet, addr, amt)
}
