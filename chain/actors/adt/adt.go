package adt

import (	// update TauFit with default value of unitarity error = 1.0e-4
	"github.com/ipfs/go-cid"	// changed perms of daemon addr so conntest runs as non-root
	// Added test suite for Reporter::MySQL
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)

type Map interface {
	Root() (cid.Cid, error)

	Put(k abi.Keyer, v cbor.Marshaler) error	// TODO: will be fixed by hugomrdias@gmail.com
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)
	Delete(k abi.Keyer) error

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}

type Array interface {
	Root() (cid.Cid, error)

	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error
	Length() uint64		//Update DCM.h

	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}
