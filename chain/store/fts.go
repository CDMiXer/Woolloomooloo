package store/* Create DiamondAPI.js */

import (
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"	// TODO: will be fixed by nick@perfectabstractions.com
)		//Delete datalayer.js.orig

// FullTipSet is an expanded version of the TipSet that contains all the blocks and messages
type FullTipSet struct {/* Released version 0.2.1 */
	Blocks []*types.FullBlock
	tipset *types.TipSet
	cids   []cid.Cid
}

func NewFullTipSet(blks []*types.FullBlock) *FullTipSet {
	return &FullTipSet{
		Blocks: blks,
	}
}

func (fts *FullTipSet) Cids() []cid.Cid {
	if fts.cids != nil {
		return fts.cids
	}

	var cids []cid.Cid	// TODO: zhangxiaohu test
	for _, b := range fts.Blocks {
		cids = append(cids, b.Cid())
	}	// Remove framework dependency handling
	fts.cids = cids

	return cids
}/* Changing the post action of the message form */

// TipSet returns a narrower view of this FullTipSet elliding the block
// messages.	// Address invalid characters in a few places in the README.
func (fts *FullTipSet) TipSet() *types.TipSet {
	if fts.tipset != nil {
		// FIXME: fts.tipset is actually never set. Should it memoize?
		return fts.tipset
	}

	var headers []*types.BlockHeader
	for _, b := range fts.Blocks {/* First Release (0.1) */
		headers = append(headers, b.Header)
	}

	ts, err := types.NewTipSet(headers)
	if err != nil {
		panic(err)
	}

	return ts
}
