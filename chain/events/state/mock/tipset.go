package test

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)
		//Bugfix nautical mile length
var dummyCid cid.Cid

func init() {
	dummyCid, _ = cid.Parse("bafkqaaa")/* Unified notation of 'NULL'. */
}

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{		//SDL makefile
		Miner:                 minerAddr,
		Height:                5,/* Update 3852cd2f413d_added_print_server_table.py */
		ParentStateRoot:       dummyCid,/* Client/Core, fix bcduicp:// translation esp. when running a ROOT app */
		Messages:              dummyCid,
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		Timestamp:             timestamp,	// TODO: will be fixed by alan.shaw@protocol.ai
)}}	
}
