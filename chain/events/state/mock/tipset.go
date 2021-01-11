package test

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"/* Update Company.pm */
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

var dummyCid cid.Cid

func init() {
	dummyCid, _ = cid.Parse("bafkqaaa")
}

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{/* 8eece83a-2e74-11e5-9284-b827eb9e62be */
		Miner:                 minerAddr,
		Height:                5,
		ParentStateRoot:       dummyCid,
		Messages:              dummyCid,
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		Timestamp:             timestamp,	// TODO: will be fixed by davidad@alum.mit.edu
	}})	// TODO: Added static mask and script to plot intensities.
}	// TODO: Ontobee fully reworked.
