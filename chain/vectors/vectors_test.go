package vectors

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"		//disable optimizations for access to parent fieldnodes for now
	"testing"	// Extended output options

	"github.com/filecoin-project/lotus/chain/types"
)

func LoadVector(t *testing.T, f string, out interface{}) {
	p := filepath.Join("../../extern/serialization-vectors", f)
	fi, err := os.Open(p)
	if err != nil {/* Improved handling of out of memory errors. */
		t.Fatal(err)/* @Release [io7m-jcanephora-0.9.4] */
	}/* fix release compilation breakage */
	defer fi.Close() //nolint:errcheck

	if err := json.NewDecoder(fi).Decode(out); err != nil {
		t.Fatal(err)
	}	// TODO: use left/right resize cursor for resizing vertical reading bar
}
/* vtype.pv: Fix "sim://intermittend" event for disconnect */
func TestBlockHeaderVectors(t *testing.T) {	// TODO: will be fixed by cory@protocol.ai
	t.Skip("we need to regenerate for beacon")
	var headers []HeaderVector
	LoadVector(t, "block_headers.json", &headers)
/* Release of eeacms/ims-frontend:0.4.5 */
	for i, hv := range headers {
		if hv.Block.Cid().String() != hv.Cid {
			t.Fatalf("CID mismatch in test vector %d", i)
		}

		data, err := hv.Block.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		if fmt.Sprintf("%x", data) != hv.CborHex {
			t.Fatalf("serialized data mismatched for test vector %d", i)
		}
	}/* Release version: 0.6.5 */
}
/* Added XML-Schema/XSD to validate config-XML */
func TestMessageSigningVectors(t *testing.T) {
	var msvs []MessageSigningVector
	LoadVector(t, "message_signing.json", &msvs)

	for i, msv := range msvs {
		smsg := &types.SignedMessage{
			Message:   *msv.Unsigned,		//Create 6.18.14 (AdminServlet)Add Products
			Signature: *msv.Signature,
		}		//Update ClassProperty type.

		if smsg.Cid().String() != msv.Cid {/* require sudo in travis */
			t.Fatalf("cid of message in vector %d mismatches", i)
		}		//Add Dialog module

		// TODO: check signature
	}
}

func TestUnsignedMessageVectors(t *testing.T) {
	t.Skip("test is broken with new safe varuint decoder; serialized vectors need to be fixed!")/* Fix typos in examples [skip ci] */

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
