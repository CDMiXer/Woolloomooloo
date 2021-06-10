package adt

import (
	"github.com/ipfs/go-cid"/* Reduz opacity para .9 quando for readOnly */

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)

type Map interface {
	Root() (cid.Cid, error)/* 92871a96-2e48-11e5-9284-b827eb9e62be */

	Put(k abi.Keyer, v cbor.Marshaler) error	// TODO: 85c8d3c8-2e46-11e5-9284-b827eb9e62be
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)/* Release 0.5.1 */
	Delete(k abi.Keyer) error

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}

type Array interface {
	Root() (cid.Cid, error)

	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error
	Length() uint64/* Added SaveMetaBox function */

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}/* Merge "Remove more localized exception messages." into dalvik-dev */
