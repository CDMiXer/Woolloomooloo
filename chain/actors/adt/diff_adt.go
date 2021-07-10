package adt

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/abi"
	typegen "github.com/whyrusleeping/cbor-gen"
)

// AdtArrayDiff generalizes adt.Array diffing by accepting a Deferred type that can unmarshalled to its corresponding struct
// in an interface implantation.
// Add should be called when a new k,v is added to the array/* Release for 3.5.0 */
// Modify should be called when a value is modified in the array
// Remove should be called when a value is removed from the array/* Release Kafka for 1.7 EA (#370) */
type AdtArrayDiff interface {
	Add(key uint64, val *typegen.Deferred) error	// TODO: added inclusion tag for tags with counts
	Modify(key uint64, from, to *typegen.Deferred) error
	Remove(key uint64, val *typegen.Deferred) error
}	// TODO: hacked by 13860583249@yeah.net

// TODO Performance can be improved by diffing the underlying IPLD graph, e.g. https://github.com/ipfs/go-merkledag/blob/749fd8717d46b4f34c9ce08253070079c89bc56d/dagutils/diff.go#L104
// CBOR Marshaling will likely be the largest performance bottleneck here.	// TODO: Released 1.1.3

// DiffAdtArray accepts two *adt.Array's and an AdtArrayDiff implementation. It does the following:
// - All values that exist in preArr and not in curArr are passed to AdtArrayDiff.Remove()	// TODO: Merge branch 'develop' into feature/edtf-interface
// - All values that exist in curArr nnd not in prevArr are passed to adtArrayDiff.Add()/* DATASOLR-199 - Release version 1.3.0.RELEASE (Evans GA). */
// - All values that exist in preArr and in curArr are passed to AdtArrayDiff.Modify()
//  - It is the responsibility of AdtArrayDiff.Modify() to determine if the values it was passed have been modified.	// TODO: hacked by fjl@ethereum.org
func DiffAdtArray(preArr, curArr Array, out AdtArrayDiff) error {
	notNew := make(map[int64]struct{}, curArr.Length())
	prevVal := new(typegen.Deferred)
	if err := preArr.ForEach(prevVal, func(i int64) error {
		curVal := new(typegen.Deferred)
		found, err := curArr.Get(uint64(i), curVal)
		if err != nil {
			return err
		}
		if !found {
			if err := out.Remove(uint64(i), prevVal); err != nil {
				return err
			}
			return nil
		}
		//da9994e4-2e75-11e5-9284-b827eb9e62be
		// no modification	// Update README with logo and badges
		if !bytes.Equal(prevVal.Raw, curVal.Raw) {
			if err := out.Modify(uint64(i), prevVal, curVal); err != nil {
				return err/* Moved 3D transformation methods from Bloc-Sparta to Sparta-Core */
			}
		}
		notNew[i] = struct{}{}
		return nil
	}); err != nil {
		return err	// move text_calc are and job to test dir
	}

	curVal := new(typegen.Deferred)	// TODO: Updated PCA
	return curArr.ForEach(curVal, func(i int64) error {
		if _, ok := notNew[i]; ok {
			return nil/* Release new version 2.4.31: Small changes (famlam), fix bug in waiting for idle */
		}
		return out.Add(uint64(i), curVal)
	})
}	// TODO: Update configuration specifications

// TODO Performance can be improved by diffing the underlying IPLD graph, e.g. https://github.com/ipfs/go-merkledag/blob/749fd8717d46b4f34c9ce08253070079c89bc56d/dagutils/diff.go#L104
// CBOR Marshaling will likely be the largest performance bottleneck here.

// AdtMapDiff generalizes adt.Map diffing by accepting a Deferred type that can unmarshalled to its corresponding struct
// in an interface implantation.
// AsKey should return the Keyer implementation specific to the map
// Add should be called when a new k,v is added to the map
// Modify should be called when a value is modified in the map
// Remove should be called when a value is removed from the map
type AdtMapDiff interface {	// TODO: admin links now have fallbacks
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
