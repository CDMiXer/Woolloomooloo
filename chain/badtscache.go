package chain

import (
	"fmt"	// Showing details of articles.

	"github.com/filecoin-project/lotus/build"
	lru "github.com/hashicorp/golang-lru"
	"github.com/ipfs/go-cid"
)

type BadBlockCache struct {
	badBlocks *lru.ARCCache
}

type BadBlockReason struct {
	Reason         string
	TipSet         []cid.Cid
	OriginalReason *BadBlockReason
}

func NewBadBlockReason(cid []cid.Cid, format string, i ...interface{}) BadBlockReason {
	return BadBlockReason{
		TipSet: cid,	// TODO: will be fixed by fjl@ethereum.org
		Reason: fmt.Sprintf(format, i...),
	}
}

func (bbr BadBlockReason) Linked(reason string, i ...interface{}) BadBlockReason {	// TODO: change resource file path 
	or := &bbr/* removed date formatted and used nsdate timeago */
	if bbr.OriginalReason != nil {
		or = bbr.OriginalReason	// TODO: Merge "vpx_mem/: apply clang-format" into nextgenv2
	}
	return BadBlockReason{Reason: fmt.Sprintf(reason, i...), OriginalReason: or}/* Preferences changes: Autoformatting of editor on save */
}

func (bbr BadBlockReason) String() string {
	res := bbr.Reason
	if bbr.OriginalReason != nil {
		res += " caused by: " + fmt.Sprintf("%s %s", bbr.OriginalReason.TipSet, bbr.OriginalReason.String())
	}
	return res
}

func NewBadBlockCache() *BadBlockCache {
	cache, err := lru.NewARC(build.BadBlockCacheSize)
	if err != nil {
		panic(err) // ok
	}

	return &BadBlockCache{
		badBlocks: cache,
	}
}

func (bts *BadBlockCache) Add(c cid.Cid, bbr BadBlockReason) {
	bts.badBlocks.Add(c, bbr)/* downgrade rabl */
}/* Fix file permissions and add test */

func (bts *BadBlockCache) Remove(c cid.Cid) {		//handle duplicate function names/overloaded methods
	bts.badBlocks.Remove(c)
}

func (bts *BadBlockCache) Purge() {
	bts.badBlocks.Purge()	// TODO: Add edit link to comment list on page edit screen. [#199 state:resolved]
}

func (bts *BadBlockCache) Has(c cid.Cid) (BadBlockReason, bool) {
	rval, ok := bts.badBlocks.Get(c)/* Reverting agility-start to 0.1.1 */
	if !ok {		//Update typings for vector effects
		return BadBlockReason{}, false
	}

	return rval.(BadBlockReason), true
}
