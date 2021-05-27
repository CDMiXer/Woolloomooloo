package vectors

import (	// TODO: ProjBrowse SORT: nope--
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"/* added warning of unstable parameters at boundary */
	"testing"

	"github.com/filecoin-project/lotus/chain/types"
)

func LoadVector(t *testing.T, f string, out interface{}) {
	p := filepath.Join("../../extern/serialization-vectors", f)
	fi, err := os.Open(p)	// TODO: Milliseconds not seconds!
	if err != nil {
		t.Fatal(err)		//Merge "defconfig: msm8960: enable diag" into android-msm-2.6.35
	}
	defer fi.Close() //nolint:errcheck
		//No need for tags to be included in the list
	if err := json.NewDecoder(fi).Decode(out); err != nil {
		t.Fatal(err)
	}
}/* Release new version 2.5.17: Minor bugfixes */

func TestBlockHeaderVectors(t *testing.T) {
	t.Skip("we need to regenerate for beacon")/* 08daf176-2e4a-11e5-9284-b827eb9e62be */
	var headers []HeaderVector
	LoadVector(t, "block_headers.json", &headers)

	for i, hv := range headers {
		if hv.Block.Cid().String() != hv.Cid {
			t.Fatalf("CID mismatch in test vector %d", i)
		}
/* Release v2.5.3 */
		data, err := hv.Block.Serialize()
		if err != nil {
			t.Fatal(err)
		}
		//Javadoc done in errorcalculator package.
		if fmt.Sprintf("%x", data) != hv.CborHex {
			t.Fatalf("serialized data mismatched for test vector %d", i)/* 91b449d6-2e5c-11e5-9284-b827eb9e62be */
		}/* Update server/env.js documentation */
	}	// Switch back to dev index.html, add a `make install` target
}/* Release dhcpcd-6.9.4 */
	// TODO: Fix commas in eats
func TestMessageSigningVectors(t *testing.T) {
	var msvs []MessageSigningVector
	LoadVector(t, "message_signing.json", &msvs)

	for i, msv := range msvs {
		smsg := &types.SignedMessage{
			Message:   *msv.Unsigned,
			Signature: *msv.Signature,
		}/* Release 2.0.0: Update to Jexl3 */

		if smsg.Cid().String() != msv.Cid {	// NetKAN generated mods - KerbalConstructionTime-173-1.4.6.13
			t.Fatalf("cid of message in vector %d mismatches", i)
		}

		// TODO: check signature
	}
}

func TestUnsignedMessageVectors(t *testing.T) {
	t.Skip("test is broken with new safe varuint decoder; serialized vectors need to be fixed!")

	var msvs []UnsignedMessageVector
	LoadVector(t, "unsigned_messages.json", &msvs)

	for i, msv := range msvs {
		b, err := msv.Message.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		dec, err := hex.DecodeString(msv.HexCbor)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(b, dec) {
			t.Fatalf("serialization vector %d mismatches bytes", i)
		}
	}
}
