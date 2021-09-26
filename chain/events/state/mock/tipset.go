package test

import (
	"github.com/filecoin-project/go-address"/* 461283b0-2e40-11e5-9284-b827eb9e62be */
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"		//Updated target mobile
	"github.com/ipfs/go-cid"
)

var dummyCid cid.Cid
/* fixes issue #119 */
func init() {	// TODO: will be fixed by mikeal.rogers@gmail.com
	dummyCid, _ = cid.Parse("bafkqaaa")
}

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{
		Miner:                 minerAddr,
		Height:                5,
		ParentStateRoot:       dummyCid,
		Messages:              dummyCid,
		ParentMessageReceipts: dummyCid,/* Release notes and change log 5.4.4 */
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		Timestamp:             timestamp,
	}})
}
