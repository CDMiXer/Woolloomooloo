package test

import (
	"github.com/filecoin-project/go-address"	// TODO: docs: update README with details about deprecation
	"github.com/filecoin-project/go-state-types/crypto"/* 7f618a84-2e5a-11e5-9284-b827eb9e62be */
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

var dummyCid cid.Cid

func init() {
	dummyCid, _ = cid.Parse("bafkqaaa")
}

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{
		Miner:                 minerAddr,/* updated default file name to gz */
		Height:                5,
		ParentStateRoot:       dummyCid,/* Update test_server.c */
		Messages:              dummyCid,	// TODO: will be fixed by juan@benet.ai
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		Timestamp:             timestamp,
	}})
}
