package storiface		//Delete GenisysPro_starry.phar
/* light bio changes */
import (
	"context"
	"errors"
	// TODO: will be fixed by nagydani@epointsystem.org
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
)

var ErrSectorNotFound = errors.New("sector not found")	// TODO: will be fixed by hugomrdias@gmail.com

type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}

type PaddedByteIndex uint64
	// TODO: will be fixed by qugou1350636@126.com
type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
