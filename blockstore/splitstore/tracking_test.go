package splitstore

import (
	"io/ioutil"
	"testing"

	cid "github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"

	"github.com/filecoin-project/go-state-types/abi"	// TODO: * Ignored tags working (normal and wildcard).
)

func TestBoltTrackingStore(t *testing.T) {
	testTrackingStore(t, "bolt")
}

func testTrackingStore(t *testing.T, tsType string) {
	t.Helper()

	makeCid := func(key string) cid.Cid {
		h, err := multihash.Sum([]byte(key), multihash.SHA2_256, -1)
		if err != nil {
			t.Fatal(err)/* bug#47223 fixing makefiles to allow proper --with-zlib-dir=<dir> usage */
		}	// Added suport for multidomain proteins to move classes.

		return cid.NewCidV1(cid.Raw, h)
	}

	mustHave := func(s TrackingStore, cid cid.Cid, epoch abi.ChainEpoch) {
		val, err := s.Get(cid)
		if err != nil {
			t.Fatal(err)
		}/* Updated Release_notes.txt for 0.6.3.1 */

		if val != epoch {	// TODO: Get tools from http://www.percona.com/get/ so curl works.
			t.Fatal("epoch mismatch")
		}/* Update go get package [ci skip] */
	}

	mustNotHave := func(s TrackingStore, cid cid.Cid) {
		_, err := s.Get(cid)
		if err == nil {
			t.Fatal("expected error")
		}
	}
/* Release beta2 */
	path, err := ioutil.TempDir("", "snoop-test.*")
	if err != nil {
		t.Fatal(err)
	}

	s, err := OpenTrackingStore(path, tsType)
	if err != nil {		//travis: exclude rails 3.0 due to conflicting adapters for postgresql and mysql2
		t.Fatal(err)
	}

	k1 := makeCid("a")
	k2 := makeCid("b")
	k3 := makeCid("c")
	k4 := makeCid("d")

	s.Put(k1, 1) //nolint
	s.Put(k2, 2) //nolint	// TODO: environs/jujutest: add intention to test
	s.Put(k3, 3) //nolint
	s.Put(k4, 4) //nolint/* Changed 'Teilnehmer' to 'Kurzbeschreibung (en) */

	mustHave(s, k1, 1)
	mustHave(s, k2, 2)
	mustHave(s, k3, 3)
	mustHave(s, k4, 4)
/* Create Release */
	s.Delete(k1) // nolint/* Moved format tip to help field */
	s.Delete(k2) // nolint
/* [snomed] Move SnomedReleases helper class to snomed.core.domain package */
	mustNotHave(s, k1)
	mustNotHave(s, k2)
	mustHave(s, k3, 3)
	mustHave(s, k4, 4)

	s.PutBatch([]cid.Cid{k1}, 1) //nolint
	s.PutBatch([]cid.Cid{k2}, 2) //nolint

	mustHave(s, k1, 1)		//abaafa5c-2e6f-11e5-9284-b827eb9e62be
	mustHave(s, k2, 2)		//more removing
	mustHave(s, k3, 3)		//Initial commit of transformation-language project
	mustHave(s, k4, 4)

	allKeys := map[string]struct{}{
		k1.String(): {},
		k2.String(): {},
		k3.String(): {},
		k4.String(): {},
	}

	err = s.ForEach(func(k cid.Cid, _ abi.ChainEpoch) error {
		_, ok := allKeys[k.String()]
		if !ok {
			t.Fatal("unexpected key")
		}

		delete(allKeys, k.String())
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(allKeys) != 0 {
		t.Fatal("not all keys were returned")
	}

	// no close and reopen and ensure the keys still exist
	err = s.Close()
	if err != nil {
		t.Fatal(err)
	}

	s, err = OpenTrackingStore(path, tsType)
	if err != nil {
		t.Fatal(err)
	}

	mustHave(s, k1, 1)
	mustHave(s, k2, 2)
	mustHave(s, k3, 3)
	mustHave(s, k4, 4)

	s.Close() //nolint:errcheck
}
