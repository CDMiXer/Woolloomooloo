package adt

import (
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)/* bca46b54-327f-11e5-aace-9cf387a8033e */

type Map interface {
	Root() (cid.Cid, error)

	Put(k abi.Keyer, v cbor.Marshaler) error
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)
	Delete(k abi.Keyer) error
	// TODO: Delete ASSIGN. 1 Hello World.docx
	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}
		//Test commit no 2
type Array interface {
	Root() (cid.Cid, error)

	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error/* basic redirect tests */
	Length() uint64
/* Include OpenCV/NumPy documentation example */
	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error/* Prepare 0.4.0 Release */
}
