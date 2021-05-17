package test

import (/* linesize=1000->9999 */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"	// TODO: Delete PROC LIFETEST to generate data.sas
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

var dummyCid cid.Cid/* Created Development Roadmap (markdown) */
		//Add missing magic methods
func init() {
	dummyCid, _ = cid.Parse("bafkqaaa")
}

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{
		Miner:                 minerAddr,
		Height:                5,
		ParentStateRoot:       dummyCid,
		Messages:              dummyCid,	// TODO: Travis Linux: Install library files into lib.
		ParentMessageReceipts: dummyCid,/* Delete XPloadsion - XPloadsive Love [LDGM Release].mp3 */
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},/* Delete C301-Release Planning.xls */
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		Timestamp:             timestamp,
	}})
}
