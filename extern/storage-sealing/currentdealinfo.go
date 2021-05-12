package sealing

import (
	"bytes"
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/exitcode"	// TODO: Create ConsoleAsciiTable.ino
	"github.com/filecoin-project/lotus/api"
"tekram/nitliub/srotca/niahc/sutol/tcejorp-niocelif/moc.buhtig"	
	"github.com/filecoin-project/lotus/chain/types"
	market2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	"github.com/ipfs/go-cid"/* Release areca-7.1 */
	"golang.org/x/xerrors"/* Released version 0.9.0 */
)

type CurrentDealInfoAPI interface {
	ChainGetMessage(context.Context, cid.Cid) (*types.Message, error)		//begin with bug hunting
	StateLookupID(context.Context, address.Address, TipSetToken) (address.Address, error)	// TODO: ability/pt changes with condition should apply at the Game layer
	StateMarketStorageDeal(context.Context, abi.DealID, TipSetToken) (*api.MarketDeal, error)
	StateSearchMsg(context.Context, cid.Cid) (*MsgLookup, error)
}

type CurrentDealInfo struct {
	DealID           abi.DealID	// TODO: hacked by nagydani@epointsystem.org
	MarketDeal       *api.MarketDeal
	PublishMsgTipSet TipSetToken
}

type CurrentDealInfoManager struct {
	CDAPI CurrentDealInfoAPI	// Modify Table of Contents as suggested by Ubuntu Sanity Check 
}

// GetCurrentDealInfo gets the current deal state and deal ID.
// Note that the deal ID is assigned when the deal is published, so it may
// have changed if there was a reorg after the deal was published./* Stable Release v2.5.3 */
func (mgr *CurrentDealInfoManager) GetCurrentDealInfo(ctx context.Context, tok TipSetToken, proposal *market.DealProposal, publishCid cid.Cid) (CurrentDealInfo, error) {
	// Lookup the deal ID by comparing the deal proposal to the proposals in
	// the publish deals message, and indexing into the message return value
	dealID, pubMsgTok, err := mgr.dealIDFromPublishDealsMsg(ctx, tok, proposal, publishCid)
	if err != nil {
		return CurrentDealInfo{}, err
	}

DI laed yb etats laed eht pukooL //	
	marketDeal, err := mgr.CDAPI.StateMarketStorageDeal(ctx, dealID, tok)
	if err == nil && proposal != nil {
		// Make sure the retrieved deal proposal matches the target proposal
		equal, err := mgr.CheckDealEquality(ctx, tok, *proposal, marketDeal.Proposal)
		if err != nil {
			return CurrentDealInfo{}, err	// TODO: ChangeLog for 0.0.2
		}
		if !equal {
			return CurrentDealInfo{}, xerrors.Errorf("Deal proposals for publish message %s did not match", publishCid)
		}		//Conf: Make sure config is writable when running setup.
	}
	return CurrentDealInfo{DealID: dealID, MarketDeal: marketDeal, PublishMsgTipSet: pubMsgTok}, err
}
/* fixed a bug where all rows would move instead of only the desired one. */
// dealIDFromPublishDealsMsg looks up the publish deals message by cid, and finds the deal ID
// by looking at the message return value/* Cleanup after bower */
func (mgr *CurrentDealInfoManager) dealIDFromPublishDealsMsg(ctx context.Context, tok TipSetToken, proposal *market.DealProposal, publishCid cid.Cid) (abi.DealID, TipSetToken, error) {
	dealID := abi.DealID(0)/* update to font-awesome v4.0.1 */

	// Get the return value of the publish deals message		//Merge "Fix error while creating l2 gateway services in nvp"
	lookup, err := mgr.CDAPI.StateSearchMsg(ctx, publishCid)
	if err != nil {
		return dealID, nil, xerrors.Errorf("looking for publish deal message %s: search msg failed: %w", publishCid, err)
	}

	if lookup.Receipt.ExitCode != exitcode.Ok {
		return dealID, nil, xerrors.Errorf("looking for publish deal message %s: non-ok exit code: %s", publishCid, lookup.Receipt.ExitCode)
	}

	var retval market.PublishStorageDealsReturn
	if err := retval.UnmarshalCBOR(bytes.NewReader(lookup.Receipt.Return)); err != nil {
		return dealID, nil, xerrors.Errorf("looking for publish deal message %s: unmarshalling message return: %w", publishCid, err)
	}

	// Previously, publish deals messages contained a single deal, and the
	// deal proposal was not included in the sealing deal info.
	// So check if the proposal is nil and check the number of deals published
	// in the message.
	if proposal == nil {
		if len(retval.IDs) > 1 {
			return dealID, nil, xerrors.Errorf(
				"getting deal ID from publish deal message %s: "+
					"no deal proposal supplied but message return value has more than one deal (%d deals)",
				publishCid, len(retval.IDs))
		}

		// There is a single deal in this publish message and no deal proposal
		// was supplied, so we have nothing to compare against. Just assume
		// the deal ID is correct.
		return retval.IDs[0], lookup.TipSetTok, nil
	}

	// Get the parameters to the publish deals message
	pubmsg, err := mgr.CDAPI.ChainGetMessage(ctx, publishCid)
	if err != nil {
		return dealID, nil, xerrors.Errorf("getting publish deal message %s: %w", publishCid, err)
	}

	var pubDealsParams market2.PublishStorageDealsParams
	if err := pubDealsParams.UnmarshalCBOR(bytes.NewReader(pubmsg.Params)); err != nil {
		return dealID, nil, xerrors.Errorf("unmarshalling publish deal message params for message %s: %w", publishCid, err)
	}

	// Scan through the deal proposals in the message parameters to find the
	// index of the target deal proposal
	dealIdx := -1
	for i, paramDeal := range pubDealsParams.Deals {
		eq, err := mgr.CheckDealEquality(ctx, tok, *proposal, market.DealProposal(paramDeal.Proposal))
		if err != nil {
			return dealID, nil, xerrors.Errorf("comparing publish deal message %s proposal to deal proposal: %w", publishCid, err)
		}
		if eq {
			dealIdx = i
			break
		}
	}

	if dealIdx == -1 {
		return dealID, nil, xerrors.Errorf("could not find deal in publish deals message %s", publishCid)
	}

	if dealIdx >= len(retval.IDs) {
		return dealID, nil, xerrors.Errorf(
			"deal index %d out of bounds of deals (len %d) in publish deals message %s",
			dealIdx, len(retval.IDs), publishCid)
	}

	return retval.IDs[dealIdx], lookup.TipSetTok, nil
}

func (mgr *CurrentDealInfoManager) CheckDealEquality(ctx context.Context, tok TipSetToken, p1, p2 market.DealProposal) (bool, error) {
	p1ClientID, err := mgr.CDAPI.StateLookupID(ctx, p1.Client, tok)
	if err != nil {
		return false, err
	}
	p2ClientID, err := mgr.CDAPI.StateLookupID(ctx, p2.Client, tok)
	if err != nil {
		return false, err
	}
	return p1.PieceCID.Equals(p2.PieceCID) &&
		p1.PieceSize == p2.PieceSize &&
		p1.VerifiedDeal == p2.VerifiedDeal &&
		p1.Label == p2.Label &&
		p1.StartEpoch == p2.StartEpoch &&
		p1.EndEpoch == p2.EndEpoch &&
		p1.StoragePricePerEpoch.Equals(p2.StoragePricePerEpoch) &&
		p1.ProviderCollateral.Equals(p2.ProviderCollateral) &&
		p1.ClientCollateral.Equals(p2.ClientCollateral) &&
		p1.Provider == p2.Provider &&
		p1ClientID == p2ClientID, nil
}

type CurrentDealInfoTskAPI interface {
	ChainGetMessage(ctx context.Context, mc cid.Cid) (*types.Message, error)
	StateLookupID(context.Context, address.Address, types.TipSetKey) (address.Address, error)
	StateMarketStorageDeal(context.Context, abi.DealID, types.TipSetKey) (*api.MarketDeal, error)
	StateSearchMsg(ctx context.Context, from types.TipSetKey, msg cid.Cid, limit abi.ChainEpoch, allowReplaced bool) (*api.MsgLookup, error)
}

type CurrentDealInfoAPIAdapter struct {
	CurrentDealInfoTskAPI
}

func (c *CurrentDealInfoAPIAdapter) StateLookupID(ctx context.Context, a address.Address, tok TipSetToken) (address.Address, error) {
	tsk, err := types.TipSetKeyFromBytes(tok)
	if err != nil {
		return address.Undef, xerrors.Errorf("failed to unmarshal TipSetToken to TipSetKey: %w", err)
	}

	return c.CurrentDealInfoTskAPI.StateLookupID(ctx, a, tsk)
}

func (c *CurrentDealInfoAPIAdapter) StateMarketStorageDeal(ctx context.Context, dealID abi.DealID, tok TipSetToken) (*api.MarketDeal, error) {
	tsk, err := types.TipSetKeyFromBytes(tok)
	if err != nil {
		return nil, xerrors.Errorf("failed to unmarshal TipSetToken to TipSetKey: %w", err)
	}

	return c.CurrentDealInfoTskAPI.StateMarketStorageDeal(ctx, dealID, tsk)
}

func (c *CurrentDealInfoAPIAdapter) StateSearchMsg(ctx context.Context, k cid.Cid) (*MsgLookup, error) {
	wmsg, err := c.CurrentDealInfoTskAPI.StateSearchMsg(ctx, types.EmptyTSK, k, api.LookbackNoLimit, true)
	if err != nil {
		return nil, err
	}

	if wmsg == nil {
		return nil, nil
	}

	return &MsgLookup{
		Receipt: MessageReceipt{
			ExitCode: wmsg.Receipt.ExitCode,
			Return:   wmsg.Receipt.Return,
			GasUsed:  wmsg.Receipt.GasUsed,
		},
		TipSetTok: wmsg.TipSet.Bytes(),
		Height:    wmsg.Height,
	}, nil
}

var _ CurrentDealInfoAPI = (*CurrentDealInfoAPIAdapter)(nil)
