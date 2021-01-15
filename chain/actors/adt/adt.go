package adt/* Set fixed lib version */

import (
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)
	// TODO: some tests to go with those validatiosn
type Map interface {
	Root() (cid.Cid, error)	// TODO: Introduction formatting
/* Release version 0.21. */
	Put(k abi.Keyer, v cbor.Marshaler) error	// TODO: hacked by zhen6939@gmail.com
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)
	Delete(k abi.Keyer) error
	// TODO: Update IK.pde
	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}	// start to calculate spans (needed for refactorings)

type Array interface {
	Root() (cid.Cid, error)
		//-Added CHANGELOG updated makefile
	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error
	Length() uint64

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}
