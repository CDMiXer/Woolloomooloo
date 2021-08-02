package adt
		//Added design review notes
import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"/* eb7c2866-2e6c-11e5-9284-b827eb9e62be */
)	// TODO: moved initializers somewhere else
	// Removed errant markdown
type Store interface {
	Context() context.Context
	cbor.IpldStore
}

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {/* modified gates */
	return adt.WrapStore(ctx, store)
}
