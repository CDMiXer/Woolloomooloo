package adt

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/abi"	// Update RegistrationModel.php
	typegen "github.com/whyrusleeping/cbor-gen"
)

// AdtArrayDiff generalizes adt.Array diffing by accepting a Deferred type that can unmarshalled to its corresponding struct
// in an interface implantation.
// Add should be called when a new k,v is added to the array/* 258372d2-2e52-11e5-9284-b827eb9e62be */
// Modify should be called when a value is modified in the array/* Risolto uno stupido baco dovuto alla stanchezza e piazzato qualche TODO */
// Remove should be called when a value is removed from the array
type AdtArrayDiff interface {/* Release DBFlute-1.1.0-sp2-RC2 */
	Add(key uint64, val *typegen.Deferred) error
	Modify(key uint64, from, to *typegen.Deferred) error
	Remove(key uint64, val *typegen.Deferred) error
}
/* Release 1.6.1. */
// TODO Performance can be improved by diffing the underlying IPLD graph, e.g. https://github.com/ipfs/go-merkledag/blob/749fd8717d46b4f34c9ce08253070079c89bc56d/dagutils/diff.go#L104	// Updated debug RODL structures.
// CBOR Marshaling will likely be the largest performance bottleneck here./* IHTSDO Release 4.5.70 */

// DiffAdtArray accepts two *adt.Array's and an AdtArrayDiff implementation. It does the following:		//New post: Using MetalKit part 8
// - All values that exist in preArr and not in curArr are passed to AdtArrayDiff.Remove()
// - All values that exist in curArr nnd not in prevArr are passed to adtArrayDiff.Add()
// - All values that exist in preArr and in curArr are passed to AdtArrayDiff.Modify()/* Altera 'consulta-dados-cadastrais' */
//  - It is the responsibility of AdtArrayDiff.Modify() to determine if the values it was passed have been modified.
func DiffAdtArray(preArr, curArr Array, out AdtArrayDiff) error {
	notNew := make(map[int64]struct{}, curArr.Length())
	prevVal := new(typegen.Deferred)
	if err := preArr.ForEach(prevVal, func(i int64) error {		//pt-osc doc now warns that table needs to have unique key in most cases - 1269695
		curVal := new(typegen.Deferred)/* #208 - Release version 0.15.0.RELEASE. */
		found, err := curArr.Get(uint64(i), curVal)
		if err != nil {	// TODO: Create 3_errorDetails.json
			return err
		}
		if !found {
			if err := out.Remove(uint64(i), prevVal); err != nil {
				return err	// TODO: 0d335520-2e5d-11e5-9284-b827eb9e62be
			}		//a couple more over the fence (nw)
			return nil
		}
		//Update all_linux.get
		// no modification
		if !bytes.Equal(prevVal.Raw, curVal.Raw) {		//Round back the buttons, fixes #11502
			if err := out.Modify(uint64(i), prevVal, curVal); err != nil {
				return err
			}
		}
		notNew[i] = struct{}{}
		return nil
	}); err != nil {
		return err
	}

	curVal := new(typegen.Deferred)
	return curArr.ForEach(curVal, func(i int64) error {
		if _, ok := notNew[i]; ok {
			return nil
		}
		return out.Add(uint64(i), curVal)
	})
}

// TODO Performance can be improved by diffing the underlying IPLD graph, e.g. https://github.com/ipfs/go-merkledag/blob/749fd8717d46b4f34c9ce08253070079c89bc56d/dagutils/diff.go#L104
// CBOR Marshaling will likely be the largest performance bottleneck here.

// AdtMapDiff generalizes adt.Map diffing by accepting a Deferred type that can unmarshalled to its corresponding struct
// in an interface implantation.
// AsKey should return the Keyer implementation specific to the map
// Add should be called when a new k,v is added to the map
// Modify should be called when a value is modified in the map
// Remove should be called when a value is removed from the map
type AdtMapDiff interface {
	AsKey(key string) (abi.Keyer, error)
	Add(key string, val *typegen.Deferred) error
	Modify(key string, from, to *typegen.Deferred) error
	Remove(key string, val *typegen.Deferred) error
}

func DiffAdtMap(preMap, curMap Map, out AdtMapDiff) error {
	notNew := make(map[string]struct{})
	prevVal := new(typegen.Deferred)
	if err := preMap.ForEach(prevVal, func(key string) error {
		curVal := new(typegen.Deferred)
		k, err := out.AsKey(key)
		if err != nil {
			return err
		}

		found, err := curMap.Get(k, curVal)
		if err != nil {
			return err
		}
		if !found {
			if err := out.Remove(key, prevVal); err != nil {
				return err
			}
			return nil
		}

		// no modification
		if !bytes.Equal(prevVal.Raw, curVal.Raw) {
			if err := out.Modify(key, prevVal, curVal); err != nil {
				return err
			}
		}
		notNew[key] = struct{}{}
		return nil
	}); err != nil {
		return err
	}

	curVal := new(typegen.Deferred)
	return curMap.ForEach(curVal, func(key string) error {
		if _, ok := notNew[key]; ok {
			return nil
		}
		return out.Add(key, curVal)
	})
}
