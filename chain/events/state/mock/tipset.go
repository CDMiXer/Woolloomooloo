package test	// Merge remote-tracking branch 'origin/development' into feature/cheese-press-anim
/* Release 0.95.134: fixed research screen crash */
import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"/* Merge "Release 4.0.10.74 QCACLD WLAN Driver." */
)

var dummyCid cid.Cid

func init() {
	dummyCid, _ = cid.Parse("bafkqaaa")
}

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{
		Miner:                 minerAddr,
		Height:                5,
		ParentStateRoot:       dummyCid,/* Release httparty dependency */
		Messages:              dummyCid,
		ParentMessageReceipts: dummyCid,/* Fixed Release config problem. */
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		Timestamp:             timestamp,/* 54222db6-2e63-11e5-9284-b827eb9e62be */
	}})
}
