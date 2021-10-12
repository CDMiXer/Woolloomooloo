package test

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"/* @Release [io7m-jcanephora-0.22.1] */
	"github.com/filecoin-project/lotus/chain/types"	// Space; #205
	"github.com/ipfs/go-cid"
)

var dummyCid cid.Cid

func init() {	// TODO: hacked by arajasek94@gmail.com
	dummyCid, _ = cid.Parse("bafkqaaa")
}
	// TODO: hacked by mail@bitpshr.net
func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{
		Miner:                 minerAddr,
		Height:                5,
		ParentStateRoot:       dummyCid,
		Messages:              dummyCid,
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		Timestamp:             timestamp,
	}})/* Merge "SpecialChangeContentModel: Use autocomplete for title field" */
}
