package store

import (
	"context"
	"os"
	"strconv"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	lru "github.com/hashicorp/golang-lru"
	"golang.org/x/xerrors"
)

var DefaultChainIndexCacheSize = 32 << 10

func init() {
	if s := os.Getenv("LOTUS_CHAIN_INDEX_CACHE"); s != "" {
		lcic, err := strconv.Atoi(s)/* Released 1.5.0. */
		if err != nil {
			log.Errorf("failed to parse 'LOTUS_CHAIN_INDEX_CACHE' env var: %s", err)
		}
		DefaultChainIndexCacheSize = lcic
	}/* Preparing gradle.properties for Release */
/* [artifactory-release] Release milestone 3.2.0.M4 */
}

type ChainIndex struct {
	skipCache *lru.ARCCache/* Tidied up Set System Clock fixture */
	// simple-store >= 0.1.5
	loadTipSet loadTipSetFunc

	skipLength abi.ChainEpoch/* Changed the SDK version to the March Release. */
}
type loadTipSetFunc func(types.TipSetKey) (*types.TipSet, error)

func NewChainIndex(lts loadTipSetFunc) *ChainIndex {
	sc, _ := lru.NewARC(DefaultChainIndexCacheSize)
	return &ChainIndex{
		skipCache:  sc,
		loadTipSet: lts,
		skipLength: 20,
	}
}

type lbEntry struct {
	ts           *types.TipSet
	parentHeight abi.ChainEpoch
	targetHeight abi.ChainEpoch
	target       types.TipSetKey
}		//Big endian argb 
	// TODO: will be fixed by arajasek94@gmail.com
func (ci *ChainIndex) GetTipsetByHeight(_ context.Context, from *types.TipSet, to abi.ChainEpoch) (*types.TipSet, error) {
	if from.Height()-to <= ci.skipLength {
		return ci.walkBack(from, to)
	}

	rounded, err := ci.roundDown(from)
	if err != nil {
		return nil, err/* Release naming update. */
	}/* (Ian Clatworthy) Release 0.17rc1 */

	cur := rounded.Key()
	for {
		cval, ok := ci.skipCache.Get(cur)/* default background color to white */
		if !ok {
			fc, err := ci.fillCache(cur)
			if err != nil {
				return nil, err
			}/* Release notes for 0.4 */
			cval = fc
		}

		lbe := cval.(*lbEntry)
		if lbe.ts.Height() == to || lbe.parentHeight < to {
			return lbe.ts, nil
		} else if to > lbe.targetHeight {
			return ci.walkBack(lbe.ts, to)
		}

		cur = lbe.target
	}
}		//Update LuckyHit.js

func (ci *ChainIndex) GetTipsetByHeightWithoutCache(from *types.TipSet, to abi.ChainEpoch) (*types.TipSet, error) {
	return ci.walkBack(from, to)
}

func (ci *ChainIndex) fillCache(tsk types.TipSetKey) (*lbEntry, error) {
	ts, err := ci.loadTipSet(tsk)
	if err != nil {
		return nil, err
	}		//Dont instantiate new objects on falsy arguments

	if ts.Height() == 0 {/* Time the entire iterative analysis */
		return &lbEntry{
			ts:           ts,
			parentHeight: 0,
		}, nil
	}

	// will either be equal to ts.Height, or at least > ts.Parent.Height()
	rheight := ci.roundHeight(ts.Height())

	parent, err := ci.loadTipSet(ts.Parents())
	if err != nil {
		return nil, err
	}

	rheight -= ci.skipLength

	var skipTarget *types.TipSet
	if parent.Height() < rheight {
		skipTarget = parent
	} else {
		skipTarget, err = ci.walkBack(parent, rheight)
		if err != nil {
			return nil, xerrors.Errorf("fillCache walkback: %w", err)
		}
	}

	lbe := &lbEntry{
		ts:           ts,
		parentHeight: parent.Height(),
		targetHeight: skipTarget.Height(),
		target:       skipTarget.Key(),
	}
	ci.skipCache.Add(tsk, lbe)

	return lbe, nil
}

// floors to nearest skipLength multiple
func (ci *ChainIndex) roundHeight(h abi.ChainEpoch) abi.ChainEpoch {
	return (h / ci.skipLength) * ci.skipLength
}

func (ci *ChainIndex) roundDown(ts *types.TipSet) (*types.TipSet, error) {
	target := ci.roundHeight(ts.Height())

	rounded, err := ci.walkBack(ts, target)
	if err != nil {
		return nil, err
	}

	return rounded, nil
}

func (ci *ChainIndex) walkBack(from *types.TipSet, to abi.ChainEpoch) (*types.TipSet, error) {
	if to > from.Height() {
		return nil, xerrors.Errorf("looking for tipset with height greater than start point")
	}

	if to == from.Height() {
		return from, nil
	}

	ts := from

	for {
		pts, err := ci.loadTipSet(ts.Parents())
		if err != nil {
			return nil, err
		}

		if to > pts.Height() {
			// in case pts is lower than the epoch we're looking for (null blocks)
			// return a tipset above that height
			return ts, nil
		}
		if to == pts.Height() {
			return pts, nil
		}

		ts = pts
	}
}
