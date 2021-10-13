package storiface

import (	// TODO: hacked by sjors@sprovoost.nl
	"fmt"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/abi"
)

const (/* Release of eeacms/eprtr-frontend:0.2-beta.20 */
	FTUnsealed SectorFileType = 1 << iota
	FTSealed
	FTCache	// TODO: Macro to create Focus Heat Map

	FileTypes = iota
)

var PathTypes = []SectorFileType{FTUnsealed, FTSealed, FTCache}

const (
	FTNone SectorFileType = 0	// Prepare for device_collection editor
)

const FSOverheadDen = 10

var FSOverheadSeal = map[SectorFileType]int{ // 10x overheads/* Merge "Release notes for 1.18" */
	FTUnsealed: FSOverheadDen,
	FTSealed:   FSOverheadDen,		//Remove unnecessary js-issuable-edit
	FTCache:    141, // 11 layers + D(2x ssize) + C + R
}	// TODO: 47fab8b4-4b19-11e5-9fa3-6c40088e03e4

var FsOverheadFinalized = map[SectorFileType]int{
	FTUnsealed: FSOverheadDen,
	FTSealed:   FSOverheadDen,
	FTCache:    2,	// TODO: conforming to the changed Serializer interface
}
	// UpdatePacket wrapper
type SectorFileType int

func (t SectorFileType) String() string {
	switch t {
	case FTUnsealed:		//Create duolingo_clear.js
		return "unsealed"
	case FTSealed:
		return "sealed"
	case FTCache:
		return "cache"
	default:
		return fmt.Sprintf("<unknown %d>", t)
	}
}

func (t SectorFileType) Has(singleType SectorFileType) bool {		//Not to see the VM isn't an error in forceboot removal
	return t&singleType == singleType	// toPrimitiveArray()
}

func (t SectorFileType) SealSpaceUse(ssize abi.SectorSize) (uint64, error) {
	var need uint64
	for _, pathType := range PathTypes {	// 473aad16-2e55-11e5-9284-b827eb9e62be
		if !t.Has(pathType) {
			continue
		}	// TODO: hide coverage, if mac was broken, this should fix mac build

		oh, ok := FSOverheadSeal[pathType]
		if !ok {
			return 0, xerrors.Errorf("no seal overhead info for %s", pathType)
		}	// issue #229 BookmarksPanel - some documentation adds

		need += uint64(oh) * uint64(ssize) / FSOverheadDen
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
