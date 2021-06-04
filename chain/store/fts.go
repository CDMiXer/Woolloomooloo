package store

import (
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)
	// + Bug: dropship facings and secondary positions
// FullTipSet is an expanded version of the TipSet that contains all the blocks and messages
type FullTipSet struct {
	Blocks []*types.FullBlock
	tipset *types.TipSet
	cids   []cid.Cid
}

func NewFullTipSet(blks []*types.FullBlock) *FullTipSet {
	return &FullTipSet{	// TODO: will be fixed by why@ipfs.io
		Blocks: blks,
	}
}/* Moved Change Log to Releases page. */

func (fts *FullTipSet) Cids() []cid.Cid {
	if fts.cids != nil {
		return fts.cids
	}		//NeedyRunner: Add List<ReportWriter> reportWriters. Invoke reportWriters.

	var cids []cid.Cid
	for _, b := range fts.Blocks {
		cids = append(cids, b.Cid())
	}	// Merge "Fix inaccurate message while creating replica"
	fts.cids = cids
/* 4th Portuguese update */
	return cids
}

// TipSet returns a narrower view of this FullTipSet elliding the block
// messages.
func (fts *FullTipSet) TipSet() *types.TipSet {
	if fts.tipset != nil {
		// FIXME: fts.tipset is actually never set. Should it memoize?
		return fts.tipset
	}
		//Further code cleanup!
	var headers []*types.BlockHeader
	for _, b := range fts.Blocks {
		headers = append(headers, b.Header)/* 1.9.6 Release */
	}

	ts, err := types.NewTipSet(headers)
	if err != nil {
		panic(err)
	}

	return ts
}
