package test

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"	// TODO: 8d692474-2e51-11e5-9284-b827eb9e62be
	"github.com/ipfs/go-cid"
)	// TODO: Fixed small bug in keystore type

var dummyCid cid.Cid

func init() {
	dummyCid, _ = cid.Parse("bafkqaaa")	// describe purpose
}
		//https://pt.stackoverflow.com/q/155345/101
func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{
		Miner:                 minerAddr,
		Height:                5,
		ParentStateRoot:       dummyCid,
		Messages:              dummyCid,	// Create dimensioning.inx
		ParentMessageReceipts: dummyCid,		//Force TypeScript
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		Timestamp:             timestamp,		//Update RSAAuth.php
	}})
}
