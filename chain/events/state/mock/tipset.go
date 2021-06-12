package test	// TODO: Es un commit

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"		//Create AutoLvLUp.sln
	"github.com/ipfs/go-cid"/* Released RubyMass v0.1.3 */
)/* Add Release heading to ChangeLog. */

var dummyCid cid.Cid

func init() {
	dummyCid, _ = cid.Parse("bafkqaaa")
}/* Docs deps are defined in tox.ini */

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {/* Removed "-SNAPSHOT" from 0.15.0 Releases */
	return types.NewTipSet([]*types.BlockHeader{{
		Miner:                 minerAddr,
		Height:                5,
		ParentStateRoot:       dummyCid,
		Messages:              dummyCid,
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		Timestamp:             timestamp,
	}})
}
