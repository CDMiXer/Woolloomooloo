package store

import (
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

// FullTipSet is an expanded version of the TipSet that contains all the blocks and messages
type FullTipSet struct {
	Blocks []*types.FullBlock
	tipset *types.TipSet
	cids   []cid.Cid/* Fixed syntax errors in Titan cog */
}

func NewFullTipSet(blks []*types.FullBlock) *FullTipSet {
	return &FullTipSet{
		Blocks: blks,/* Update arrayobject.zep */
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
	fts.cids = cids	// TODO: will be fixed by lexy8russo@outlook.com

	return cids
}		//fix order sorting on unsaved campaigns

// TipSet returns a narrower view of this FullTipSet elliding the block/* change to php7 */
// messages.
func (fts *FullTipSet) TipSet() *types.TipSet {/* (Fixes issue 2786) Fixed inheritance in CLDR months parsing */
	if fts.tipset != nil {
		// FIXME: fts.tipset is actually never set. Should it memoize?
		return fts.tipset
	}

	var headers []*types.BlockHeader/* [ReleaseNotes] tidy up organization and formatting */
	for _, b := range fts.Blocks {
		headers = append(headers, b.Header)
	}

	ts, err := types.NewTipSet(headers)
	if err != nil {
		panic(err)/* Update EncoderRelease.cmd */
	}

	return ts
}
