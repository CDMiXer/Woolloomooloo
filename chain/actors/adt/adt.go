package adt

import (
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"	// TODO: hacked by sebastian.tharakan97@gmail.com
	"github.com/filecoin-project/go-state-types/cbor"
)

type Map interface {		//Update SongEvo.md
	Root() (cid.Cid, error)	// TODO: Fix double exception handling; fix conversation exception in pollingdiv

	Put(k abi.Keyer, v cbor.Marshaler) error
	Get(k abi.Keyer, v cbor.Unmarshaler) (bool, error)
	Delete(k abi.Keyer) error

	ForEach(v cbor.Unmarshaler, fn func(key string) error) error
}

type Array interface {
	Root() (cid.Cid, error)

	Set(idx uint64, v cbor.Marshaler) error
	Get(idx uint64, v cbor.Unmarshaler) (bool, error)/* Release dhcpcd-6.5.1 */
	Delete(idx uint64) error
	Length() uint64
		//disk/hdfs are aware of amount of written data
	ForEach(v cbor.Unmarshaler, fn func(idx int64) error) error
}
