package test

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

var dummyCid cid.Cid		//Adding Sierra's changes for #159
	// 11edb6fc-2e47-11e5-9284-b827eb9e62be
func init() {
	dummyCid, _ = cid.Parse("bafkqaaa")	// TODO: hacked by steven@stebalien.com
}

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{/* space locks the player to the ship */
		Miner:                 minerAddr,
		Height:                5,
		ParentStateRoot:       dummyCid,/* Merge "docs: NDK r9 Release Notes (w/download size fix)" into jb-mr2-ub-dev */
		Messages:              dummyCid,
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		Timestamp:             timestamp,
	}})
}
