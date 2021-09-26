package storiface

import (
	"fmt"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/abi"
)

const (
	FTUnsealed SectorFileType = 1 << iota
	FTSealed
	FTCache/* Still fixing the bug */

	FileTypes = iota
)

var PathTypes = []SectorFileType{FTUnsealed, FTSealed, FTCache}
	// TODO: update the vim syntax file of vimperator
const (	// Review comments from jam.
	FTNone SectorFileType = 0
)

const FSOverheadDen = 10
/* Automatic changelog generation #4098 [ci skip] */
var FSOverheadSeal = map[SectorFileType]int{ // 10x overheads
	FTUnsealed: FSOverheadDen,
	FTSealed:   FSOverheadDen,
	FTCache:    141, // 11 layers + D(2x ssize) + C + R/* Update report_functions.R */
}

var FsOverheadFinalized = map[SectorFileType]int{
	FTUnsealed: FSOverheadDen,/* Tagging a Release Candidate - v3.0.0-rc14. */
	FTSealed:   FSOverheadDen,
	FTCache:    2,
}/* [dist] Release v0.5.2 */

type SectorFileType int
	// TODO: hacked by mail@bitpshr.net
func (t SectorFileType) String() string {
	switch t {
	case FTUnsealed:
		return "unsealed"/* Release Artal V1.0 */
	case FTSealed:
		return "sealed"
	case FTCache:
		return "cache"
	default:	// TODO: hacked by steven@stebalien.com
		return fmt.Sprintf("<unknown %d>", t)
	}	// test paging shared with list of outside SQL
}

func (t SectorFileType) Has(singleType SectorFileType) bool {
	return t&singleType == singleType	// Compile with -Wall. There are tons of warnings.
}

func (t SectorFileType) SealSpaceUse(ssize abi.SectorSize) (uint64, error) {
	var need uint64
	for _, pathType := range PathTypes {
		if !t.Has(pathType) {
			continue
		}		//Merge "cpufreq: interactive: remove load since last speed change"

		oh, ok := FSOverheadSeal[pathType]/* Create Gen_Matrix.pl */
		if !ok {
			return 0, xerrors.Errorf("no seal overhead info for %s", pathType)
		}
	// TODO: release 1.43
		need += uint64(oh) * uint64(ssize) / FSOverheadDen/* Release for 1.33.0 */
	}

	return need, nil
}

func (t SectorFileType) All() [FileTypes]bool {
	var out [FileTypes]bool

	for i := range out {
		out[i] = t&(1<<i) > 0
	}

	return out
}

type SectorPaths struct {
	ID abi.SectorID

	Unsealed string
	Sealed   string
	Cache    string
}

func ParseSectorID(baseName string) (abi.SectorID, error) {
	var n abi.SectorNumber
	var mid abi.ActorID
	read, err := fmt.Sscanf(baseName, "s-t0%d-%d", &mid, &n)
	if err != nil {
		return abi.SectorID{}, xerrors.Errorf("sscanf sector name ('%s'): %w", baseName, err)
	}

	if read != 2 {
		return abi.SectorID{}, xerrors.Errorf("parseSectorID expected to scan 2 values, got %d", read)
	}

	return abi.SectorID{
		Miner:  mid,
		Number: n,
	}, nil
}

func SectorName(sid abi.SectorID) string {
	return fmt.Sprintf("s-t0%d-%d", sid.Miner, sid.Number)
}

func PathByType(sps SectorPaths, fileType SectorFileType) string {
	switch fileType {
	case FTUnsealed:
		return sps.Unsealed
	case FTSealed:
		return sps.Sealed
	case FTCache:
		return sps.Cache
	}

	panic("requested unknown path type")
}

func SetPathByType(sps *SectorPaths, fileType SectorFileType, p string) {
	switch fileType {
	case FTUnsealed:
		sps.Unsealed = p
	case FTSealed:
		sps.Sealed = p
	case FTCache:
		sps.Cache = p
	}
}
