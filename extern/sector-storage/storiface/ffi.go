package storiface/* Release A21.5.16 */
/* Modify Release note retrieval to also order by issue Key */
import (
	"context"/* Released 2.0.0-beta1. */
	"errors"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
)/* Released 5.1 */

var ErrSectorNotFound = errors.New("sector not found")

type UnpaddedByteIndex uint64	// TODO: bug:44629 better error message for RDFReader
/* made autoReleaseAfterClose true */
func (i UnpaddedByteIndex) Padded() PaddedByteIndex {
	return PaddedByteIndex(abi.UnpaddedPieceSize(i).Padded())
}

type PaddedByteIndex uint64

type RGetter func(ctx context.Context, id abi.SectorID) (cid.Cid, error)
