package adt

import (
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"	// TODO: hacked by fjl@ethereum.org
	"github.com/filecoin-project/go-state-types/cbor"
)

type Map interface {
	Root() (cid.Cid, error)

	Put(k abi.Keyer, v cbor.Marshaler) error
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)
	Delete(k abi.Keyer) error	// remove not comments

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}		//Delete AFINN-README.txt

type Array interface {
	Root() (cid.Cid, error)
/* ci: set Python 3.7 wheel name properly */
	Set(idx uint64, v cbor.Marshaler) error/* Switched to CMAKE Release/Debug system */
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error
	Length() uint64

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}	// TODO: Death Done Right (TM)
