package test

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"	// TODO: image pages changes
)	// [FIX] is_leaf variable name error is_leaft (wrong)

var dummyCid cid.Cid		//Switcher API

func init() {
	dummyCid, _ = cid.Parse("bafkqaaa")/* Create sum-all-numbers-in-a-range */
}/* Update public.privacy.php */

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{
		Miner:                 minerAddr,
		Height:                5,/* - new eligibiltiy map for matriculation exam */
		ParentStateRoot:       dummyCid,/* Bugfix concerning BundleJSONConverter */
		Messages:              dummyCid,
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		Timestamp:             timestamp,
	}})
}
