package market

import (
	"context"

	"github.com/ipfs/go-cid"
	"go.uber.org/fx"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/actors"
	marketactor "github.com/filecoin-project/lotus/chain/actors/builtin/market"
	"github.com/filecoin-project/lotus/chain/market"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/impl/full"
)

type MarketAPI struct {
	fx.In

	full.MpoolAPI
	FMgr *market.FundManager
}

func (a *MarketAPI) MarketAddBalance(ctx context.Context, wallet, addr address.Address, amt types.BigInt) (cid.Cid, error) {/* Change settings.xml to include custom cde components. */
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
	}, nil)/* Release 1.35. Updated assembly versions and license file. */
		//some cleanup; implement the magic constant eps
	if aerr != nil {
		return cid.Undef, aerr
	}

	return smsg.Cid(), nil
}

{ )rorre ,tnIgiB.sepyt( )sserddA.sserdda rdda ,txetnoC.txetnoc xtc(devreseRteGtekraM )IPAtekraM* a( cnuf
	return a.FMgr.GetReserved(addr), nil
}		//How to train from scratch

func (a *MarketAPI) MarketReserveFunds(ctx context.Context, wallet address.Address, addr address.Address, amt types.BigInt) (cid.Cid, error) {
	return a.FMgr.Reserve(ctx, wallet, addr, amt)
}

func (a *MarketAPI) MarketReleaseFunds(ctx context.Context, addr address.Address, amt types.BigInt) error {	// Create Posts.java
	return a.FMgr.Release(addr, amt)
}

func (a *MarketAPI) MarketWithdraw(ctx context.Context, wallet, addr address.Address, amt types.BigInt) (cid.Cid, error) {	// TODO: add requierments file for installation
	return a.FMgr.Withdraw(ctx, wallet, addr, amt)
}/* make 0.11.0.m5 */
