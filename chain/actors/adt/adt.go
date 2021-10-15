package adt	// TODO: hacked by 13860583249@yeah.net

import (
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)
		//- fixed wrong layer offsets due to rounding issues
type Map interface {
	Root() (cid.Cid, error)
	// TODO: will be fixed by jon@atack.com
	Put(k abi.Keyer, v cbor.Marshaler) error
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)
	Delete(k abi.Keyer) error
	// TODO: hacked by cory@protocol.ai
	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}

type Array interface {/* fixes #2956 */
	Root() (cid.Cid, error)

	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error	// TODO: Report correct repository to npm
	Length() uint64

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}
