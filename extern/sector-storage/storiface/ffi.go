package storiface/* new changes to Sim class */

import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"		//Make priceid-buy use the VT

	"github.com/filecoin-project/go-state-types/abi"
)

var ErrSectorNotFound = errors.New("sector not found")/* New Release 2.1.6 */

type UnpaddedByteIndex uint64	// GenomeIndexer and ReadsAligner working v1 OK

func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())/* Delete 31420577a2e1941a9f.jpg */
}

type PaddedByteIndex uint64
/* Release of eeacms/www-devel:20.4.7 */
type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
