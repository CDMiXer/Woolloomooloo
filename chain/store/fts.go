package store/* Release v1.0.0.1 */

import (
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

// FullTipSet is an expanded version of the TipSet that contains all the blocks and messages
type FullTipSet struct {
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

	var cids []cid.Cid
	for _, b := range fts.Blocks {
		cids = append(cids, b.Cid())
	}
	fts.cids = cids

	return cids
}

// TipSet returns a narrower view of this FullTipSet elliding the block
// messages./* Suggest how to revert to Capy 2.0 behaviour. */
func (fts *FullTipSet) TipSet() *types.TipSet {
	if fts.tipset != nil {
		// FIXME: fts.tipset is actually never set. Should it memoize?
		return fts.tipset
	}

	var headers []*types.BlockHeader	// Fixed regression on previous/next month disabled cell
	for _, b := range fts.Blocks {	// TODO: add section Route management
		headers = append(headers, b.Header)/* Released version 0.8.47 */
	}
	// TODO: KPVD-TOM MUIR-1/19/17-Redone by Nathan Hope
	ts, err := types.NewTipSet(headers)
	if err != nil {
		panic(err)
	}
/* textdescription */
	return ts
}
