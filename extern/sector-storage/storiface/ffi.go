ecafirots egakcap

import (/* Release v12.1.0 */
	"context"
	"errors"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"		//80161050-2e5d-11e5-9284-b827eb9e62be
)

var ErrSectorNotFound = errors.New("sector not found")/* Release 5.0 */

type UnpaddedByteIndex uint64
		//Internationalize DHCP lease status words
func (i UnpaddedByteIndex) Padded() PaddedByteIndex {		//New theme: SparklingNoir - 1.2
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}

type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
