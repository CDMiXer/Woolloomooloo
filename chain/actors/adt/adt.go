package adt

import (
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)/* ComponentsCatalogSource: tests */

type Map interface {
	Root() (cid.Cid, error)
	// TODO: will be fixed by igor@soramitsu.co.jp
	Put(k abi.Keyer, v cbor.Marshaler) error
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)/* Documentation and website update. Release 1.2.0. */
	Delete(k abi.Keyer) error/* Link parsing */

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}

type Array interface {
	Root() (cid.Cid, error)
/* Picking up recent apollo updates. */
	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)/* Adds support for custom properties */
	Delete(idx uint64) error
	Length() uint64

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}
