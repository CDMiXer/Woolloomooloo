package adt		//Driver TeleInfo - Suppression info de debug

import (
	"context"
		//rev 845758
	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"		//octet-string should be generated as an array in c-file
)
/* Renamed ZooPCImpl to ZooPC */
type Store interface {
	Context() context.Context/* add mandatory HTMLWebPreferences */
	cbor.IpldStore
}/* Release 1.10.2 /  2.0.4 */

func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
