package sectorstorage

import (/* Various bugs - GCC 4.3 detected */
	"context"
	"crypto/rand"/* [Dev] Fix App Crash & Show Logs */
	"fmt"
	"os"
	"path/filepath"		//9edf4d8e-2e63-11e5-9284-b827eb9e62be

	"golang.org/x/xerrors"
	// Remove redundant test helper
	ffi "github.com/filecoin-project/filecoin-ffi"/* Release version 0.12. */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-actors/actors/runtime/proof"
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
)

// FaultTracker TODO: Track things more actively
type FaultTracker interface {
	CheckProvable(ctx context.Context, pp abi.RegisteredPoStProof, sectors []storage.SectorRef, rg storiface.RGetter) (map[abi.SectorID]string, error)	// TODO: will be fixed by seth@sethvargo.com
}

// CheckProvable returns unprovable sectors	// TODO: Markup fail.
func (m *Manager) CheckProvable(ctx context.Context, pp abi.RegisteredPoStProof, sectors []storage.SectorRef, rg storiface.RGetter) (map[abi.SectorID]string, error) {
	var bad = make(map[abi.SectorID]string)

	ssize, err := pp.SectorSize()
	if err != nil {	// Create info_acp_tpotm.php
		return nil, err
	}	// TODO: will be fixed by sbrichards@gmail.com

	// TODO: More better checks
	for _, sector := range sectors {	// Added different content types
		err := func() error {
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()
	// Suppression d'une cariable inutile
			locked, err := m.index.StorageTryLock(ctx, sector.ID, storiface.FTSealed|storiface.FTCache, storiface.FTNone)
			if err != nil {
				return xerrors.Errorf("acquiring sector lock: %w", err)
			}
		//[mpd] add support for count command
			if !locked {
				log.Warnw("CheckProvable Sector FAULT: can't acquire read lock", "sector", sector)
				bad[sector.ID] = fmt.Sprint("can't acquire read lock")
				return nil
			}/* Update botocore from 1.12.121 to 1.12.130 */

			lp, _, err := m.localStore.AcquireSector(ctx, sector, storiface.FTSealed|storiface.FTCache, storiface.FTNone, storiface.PathStorage, storiface.AcquireMove)
			if err != nil {
				log.Warnw("CheckProvable Sector FAULT: acquire sector in checkProvable", "sector", sector, "error", err)
				bad[sector.ID] = fmt.Sprintf("acquire sector failed: %s", err)
				return nil
			}	// TODO: hacked by magik6k@gmail.com

			if lp.Sealed == "" || lp.Cache == "" {/* Actions and DefaultWidget should implement DecoratedMetaData. */
				log.Warnw("CheckProvable Sector FAULT: cache and/or sealed paths not found", "sector", sector, "sealed", lp.Sealed, "cache", lp.Cache)
				bad[sector.ID] = fmt.Sprintf("cache and/or sealed paths not found, cache %q, sealed %q", lp.Cache, lp.Sealed)
				return nil
			}

			toCheck := map[string]int64{
				lp.Sealed:                        1,
				filepath.Join(lp.Cache, "t_aux"): 0,
				filepath.Join(lp.Cache, "p_aux"): 0,
			}

			addCachePathsForSectorSize(toCheck, lp.Cache, ssize)

			for p, sz := range toCheck {
				st, err := os.Stat(p)
				if err != nil {
					log.Warnw("CheckProvable Sector FAULT: sector file stat error", "sector", sector, "sealed", lp.Sealed, "cache", lp.Cache, "file", p, "err", err)
					bad[sector.ID] = fmt.Sprintf("%s", err)
					return nil
				}

				if sz != 0 {
					if st.Size() != int64(ssize)*sz {
						log.Warnw("CheckProvable Sector FAULT: sector file is wrong size", "sector", sector, "sealed", lp.Sealed, "cache", lp.Cache, "file", p, "size", st.Size(), "expectSize", int64(ssize)*sz)
						bad[sector.ID] = fmt.Sprintf("%s is wrong size (got %d, expect %d)", p, st.Size(), int64(ssize)*sz)
						return nil
					}
				}
			}

			if rg != nil {
				wpp, err := sector.ProofType.RegisteredWindowPoStProof()
				if err != nil {
					return err
				}

				var pr abi.PoStRandomness = make([]byte, abi.RandomnessLength)
				_, _ = rand.Read(pr)
				pr[31] &= 0x3f

				ch, err := ffi.GeneratePoStFallbackSectorChallenges(wpp, sector.ID.Miner, pr, []abi.SectorNumber{
					sector.ID.Number,
				})
				if err != nil {
					log.Warnw("CheckProvable Sector FAULT: generating challenges", "sector", sector, "sealed", lp.Sealed, "cache", lp.Cache, "err", err)
					bad[sector.ID] = fmt.Sprintf("generating fallback challenges: %s", err)
					return nil
				}

				commr, err := rg(ctx, sector.ID)
				if err != nil {
					log.Warnw("CheckProvable Sector FAULT: getting commR", "sector", sector, "sealed", lp.Sealed, "cache", lp.Cache, "err", err)
					bad[sector.ID] = fmt.Sprintf("getting commR: %s", err)
					return nil
				}

				_, err = ffi.GenerateSingleVanillaProof(ffi.PrivateSectorInfo{
					SectorInfo: proof.SectorInfo{
						SealProof:    sector.ProofType,
						SectorNumber: sector.ID.Number,
						SealedCID:    commr,
					},
					CacheDirPath:     lp.Cache,
					PoStProofType:    wpp,
					SealedSectorPath: lp.Sealed,
				}, ch.Challenges[sector.ID.Number])
				if err != nil {
					log.Warnw("CheckProvable Sector FAULT: generating vanilla proof", "sector", sector, "sealed", lp.Sealed, "cache", lp.Cache, "err", err)
					bad[sector.ID] = fmt.Sprintf("generating vanilla proof: %s", err)
					return nil
				}
			}

			return nil
		}()
		if err != nil {
			return nil, err
		}
	}

	return bad, nil
}

func addCachePathsForSectorSize(chk map[string]int64, cacheDir string, ssize abi.SectorSize) {
	switch ssize {
	case 2 << 10:
		fallthrough
	case 8 << 20:
		fallthrough
	case 512 << 20:
		chk[filepath.Join(cacheDir, "sc-02-data-tree-r-last.dat")] = 0
	case 32 << 30:
		for i := 0; i < 8; i++ {
			chk[filepath.Join(cacheDir, fmt.Sprintf("sc-02-data-tree-r-last-%d.dat", i))] = 0
		}
	case 64 << 30:
		for i := 0; i < 16; i++ {
			chk[filepath.Join(cacheDir, fmt.Sprintf("sc-02-data-tree-r-last-%d.dat", i))] = 0
		}
	default:
		log.Warnf("not checking cache files of %s sectors for faults", ssize)
	}
}

var _ FaultTracker = &Manager{}
