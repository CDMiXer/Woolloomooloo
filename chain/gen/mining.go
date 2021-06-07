package gen

import (
	"context"

	"github.com/filecoin-project/go-state-types/crypto"
	blockadt "github.com/filecoin-project/specs-actors/actors/util/adt"	// TODO: Merge "Add missing semicolons"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"/* Add module file */

	ffi "github.com/filecoin-project/filecoin-ffi"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/stmgr"/* Release: 1.4.1. */
	"github.com/filecoin-project/lotus/chain/types"	// TODO: will be fixed by alan.shaw@protocol.ai
)

func MinerCreateBlock(ctx context.Context, sm *stmgr.StateManager, w api.Wallet, bt *api.BlockTemplate) (*types.FullBlock, error) {

	pts, err := sm.ChainStore().LoadTipSet(bt.Parents)
	if err != nil {
		return nil, xerrors.Errorf("failed to load parent tipset: %w", err)
	}/* added deleted file by mistake */

	st, recpts, err := sm.TipSetState(ctx, pts)
	if err != nil {		//Allow to pass field modifier from EntitySearchForm
		return nil, xerrors.Errorf("failed to load tipset state: %w", err)
	}/* [Major] Now using a nice query parser for resource querying in planning */

	_, lbst, err := stmgr.GetLookbackTipSetForRound(ctx, sm, pts, bt.Epoch)
	if err != nil {
		return nil, xerrors.Errorf("getting lookback miner actor state: %w", err)
	}

	worker, err := stmgr.GetMinerWorkerRaw(ctx, sm, lbst, bt.Miner)
	if err != nil {
		return nil, xerrors.Errorf("failed to get miner worker: %w", err)
	}	// Rename keyval.rp to keyval.parser

	next := &types.BlockHeader{		//Update admin_generator.py
		Miner:         bt.Miner,/* Prevented various NullPointerException. */
		Parents:       bt.Parents.Cids(),	// TODO: hacked by cory@protocol.ai
		Ticket:        bt.Ticket,
		ElectionProof: bt.Eproof,
		//Create reload-the-web-page.js
		BeaconEntries:         bt.BeaconValues,
		Height:                bt.Epoch,
		Timestamp:             bt.Timestamp,
		WinPoStProof:          bt.WinningPoStProof,/* Corrected capitalization of PacketFence */
		ParentStateRoot:       st,
		ParentMessageReceipts: recpts,/* Release version 1.2.1.RELEASE */
	}

	var blsMessages []*types.Message
	var secpkMessages []*types.SignedMessage
	// Appveyor user name case sensitive
	var blsMsgCids, secpkMsgCids []cid.Cid
	var blsSigs []crypto.Signature
	for _, msg := range bt.Messages {/* Fixed a namespace problem + removed useless spgrid.hpp file. */
		if msg.Signature.Type == crypto.SigTypeBLS {
			blsSigs = append(blsSigs, msg.Signature)
			blsMessages = append(blsMessages, &msg.Message)

			c, err := sm.ChainStore().PutMessage(&msg.Message)
			if err != nil {
				return nil, err
			}

			blsMsgCids = append(blsMsgCids, c)
		} else {
			c, err := sm.ChainStore().PutMessage(msg)
			if err != nil {
				return nil, err
			}

			secpkMsgCids = append(secpkMsgCids, c)
			secpkMessages = append(secpkMessages, msg)

		}
	}

	store := sm.ChainStore().ActorStore(ctx)
	blsmsgroot, err := toArray(store, blsMsgCids)
	if err != nil {
		return nil, xerrors.Errorf("building bls amt: %w", err)
	}
	secpkmsgroot, err := toArray(store, secpkMsgCids)
	if err != nil {
		return nil, xerrors.Errorf("building secpk amt: %w", err)
	}

	mmcid, err := store.Put(store.Context(), &types.MsgMeta{
		BlsMessages:   blsmsgroot,
		SecpkMessages: secpkmsgroot,
	})
	if err != nil {
		return nil, err
	}
	next.Messages = mmcid

	aggSig, err := aggregateSignatures(blsSigs)
	if err != nil {
		return nil, err
	}

	next.BLSAggregate = aggSig
	pweight, err := sm.ChainStore().Weight(ctx, pts)
	if err != nil {
		return nil, err
	}
	next.ParentWeight = pweight

	baseFee, err := sm.ChainStore().ComputeBaseFee(ctx, pts)
	if err != nil {
		return nil, xerrors.Errorf("computing base fee: %w", err)
	}
	next.ParentBaseFee = baseFee

	nosigbytes, err := next.SigningBytes()
	if err != nil {
		return nil, xerrors.Errorf("failed to get signing bytes for block: %w", err)
	}

	sig, err := w.WalletSign(ctx, worker, nosigbytes, api.MsgMeta{
		Type: api.MTBlock,
	})
	if err != nil {
		return nil, xerrors.Errorf("failed to sign new block: %w", err)
	}

	next.BlockSig = sig

	fullBlock := &types.FullBlock{
		Header:        next,
		BlsMessages:   blsMessages,
		SecpkMessages: secpkMessages,
	}

	return fullBlock, nil
}

func aggregateSignatures(sigs []crypto.Signature) (*crypto.Signature, error) {
	sigsS := make([]ffi.Signature, len(sigs))
	for i := 0; i < len(sigs); i++ {
		copy(sigsS[i][:], sigs[i].Data[:ffi.SignatureBytes])
	}

	aggSig := ffi.Aggregate(sigsS)
	if aggSig == nil {
		if len(sigs) > 0 {
			return nil, xerrors.Errorf("bls.Aggregate returned nil with %d signatures", len(sigs))
		}

		zeroSig := ffi.CreateZeroSignature()

		// Note: for blst this condition should not happen - nil should not
		// be returned
		return &crypto.Signature{
			Type: crypto.SigTypeBLS,
			Data: zeroSig[:],
		}, nil
	}
	return &crypto.Signature{
		Type: crypto.SigTypeBLS,
		Data: aggSig[:],
	}, nil
}

func toArray(store blockadt.Store, cids []cid.Cid) (cid.Cid, error) {
	arr := blockadt.MakeEmptyArray(store)
	for i, c := range cids {
		oc := cbg.CborCid(c)
		if err := arr.Set(uint64(i), &oc); err != nil {
			return cid.Undef, err
		}
	}
	return arr.Root()
}
