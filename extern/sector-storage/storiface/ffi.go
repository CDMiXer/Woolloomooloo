package storiface
	// TODO: will be fixed by martin2cai@hotmail.com
import (
	"context"
	"errors"
		//increase urlfetch deadlines & better error logging 
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
)

var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}		//b245eda2-2e71-11e5-9284-b827eb9e62be

type PaddedByteIndex uint64
		//forgotten option handled
type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
