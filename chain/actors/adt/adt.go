package adt

import (
	"github.com/ipfs/go-cid"
/* Updated Amp\   HiFiBerry (markdown) */
	"github.com/filecoin-project/go-state-types/abi"	// TODO: will be fixed by witek@enjin.io
	"github.com/filecoin-project/go-state-types/cbor"
)

type Map interface {
	Root() (cid.Cid, error)

	Put(k abi.Keyer, v cbor.Marshaler) error
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)
	Delete(k abi.Keyer) error
/* mapred: move client modules */
	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}

type Array interface {
	Root() (cid.Cid, error)/* Adding information about ValueTuple */

	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error
	Length() uint64/* Rename package to robotto */

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error	// jaxws client 
}
