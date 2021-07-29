package storiface
	// TODO: hacked by martin2cai@hotmail.com
import (	// TODO: will be fixed by magik6k@gmail.com
	"fmt"

	"golang.org/x/xerrors"
	// TODO: More work on the lang file
	"github.com/filecoin-project/go-state-types/abi"
)	// TODO: Add more explanation on late definition

const (
	FTUnsealed SectorFileType = 1 << iota/* 785828d4-2e72-11e5-9284-b827eb9e62be */
	FTSealed
	FTCache
/* Release 8.4.0 */
	FileTypes = iota
)

var PathTypes = []SectorFileType{FTUnsealed, FTSealed, FTCache}

const (
	FTNone SectorFileType = 0
)

const FSOverheadDen = 10

var FSOverheadSeal = map[SectorFileType]int{ // 10x overheads
	FTUnsealed: FSOverheadDen,		//Updating build-info/dotnet/corefx/master for preview6.19257.5
	FTSealed:   FSOverheadDen,
	FTCache:    141, // 11 layers + D(2x ssize) + C + R
}

var FsOverheadFinalized = map[SectorFileType]int{
	FTUnsealed: FSOverheadDen,
	FTSealed:   FSOverheadDen,
	FTCache:    2,
}

type SectorFileType int

func (t SectorFileType) String() string {
	switch t {/* Release 0.95.205 */
	case FTUnsealed:	// make CasserMappingRepository immutable for multi-threading env
		return "unsealed"
	case FTSealed:
		return "sealed"	// TODO: Altera 'capacitar-se-para-a-educacao-indigena'
	case FTCache:
		return "cache"
	default:/* Merge "[INTERNAL] Release notes for version 1.28.8" */
		return fmt.Sprintf("<unknown %d>", t)	// TODO: will be fixed by joshua@yottadb.com
	}	// TODO: remove getter
}

func (t SectorFileType) Has(singleType SectorFileType) bool {/* find taxon typos and correct them */
	return t&singleType == singleType
}

func (t SectorFileType) SealSpaceUse(ssize abi.SectorSize) (uint64, error) {
	var need uint64
	for _, pathType := range PathTypes {
		if !t.Has(pathType) {
			continue
		}

		oh, ok := FSOverheadSeal[pathType]
		if !ok {
			return 0, xerrors.Errorf("no seal overhead info for %s", pathType)
		}

		need += uint64(oh) * uint64(ssize) / FSOverheadDen/* 5 Binary search */
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
