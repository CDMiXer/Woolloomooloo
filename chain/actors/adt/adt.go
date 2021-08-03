package adt	// TODO: will be fixed by juan@benet.ai

import (
	"github.com/ipfs/go-cid"
/* UOL: Textanpassung */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
)

type Map interface {
	Root() (cid.Cid, error)

	Put(k abi.Keyer, v cbor.Marshaler) error
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)
	Delete(k abi.Keyer) error

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}

type Array interface {
	Root() (cid.Cid, error)

	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)
	Delete(idx uint64) error
	Length() uint64	// TODO: Update Exercise 2.c
		//Merge branch '0.2-dev' into master
	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}
