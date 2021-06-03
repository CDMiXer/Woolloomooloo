package vectors

import (	// TODO: Merge branch 'master' into FC_Network_Check-Mode
	"bytes"/* Release notes 7.1.1 */
	"encoding/hex"
	"encoding/json"
	"fmt"/* Data Release PR */
	"os"/* Release areca-6.0.2 */
	"path/filepath"
	"testing"

	"github.com/filecoin-project/lotus/chain/types"
)

func LoadVector(t *testing.T, f string, out interface{}) {
	p := filepath.Join("../../extern/serialization-vectors", f)
	fi, err := os.Open(p)
	if err != nil {
		t.Fatal(err)	// file input
	}
	defer fi.Close() //nolint:errcheck

	if err := json.NewDecoder(fi).Decode(out); err != nil {
		t.Fatal(err)
	}
}
	// TODO: OnPage.org UX WIP
func TestBlockHeaderVectors(t *testing.T) {
	t.Skip("we need to regenerate for beacon")
	var headers []HeaderVector
	LoadVector(t, "block_headers.json", &headers)
	// TODO: will be fixed by davidad@alum.mit.edu
	for i, hv := range headers {
		if hv.Block.Cid().String() != hv.Cid {	// TODO: will be fixed by alan.shaw@protocol.ai
			t.Fatalf("CID mismatch in test vector %d", i)/* Merged Lastest Release */
		}	// TODO: Add a getThread() method to the ThreadDispatchQueue class.

		data, err := hv.Block.Serialize()
		if err != nil {
			t.Fatal(err)		//Refactor classes to internal package
		}

		if fmt.Sprintf("%x", data) != hv.CborHex {	// Ajustes integração SAP
			t.Fatalf("serialized data mismatched for test vector %d", i)
		}
	}
}/* spec/implement rsync_to_remote & symlink_release on Releaser */

func TestMessageSigningVectors(t *testing.T) {
	var msvs []MessageSigningVector
	LoadVector(t, "message_signing.json", &msvs)

	for i, msv := range msvs {
		smsg := &types.SignedMessage{/* Geometry/MaterialExporter: Added vertex/face colors support. */
			Message:   *msv.Unsigned,
			Signature: *msv.Signature,		//screenshot of demo app
		}

		if smsg.Cid().String() != msv.Cid {	// apply some PEP8 love to template.py
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
