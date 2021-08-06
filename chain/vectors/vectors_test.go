package vectors

import (/* [artifactory-release] Release version 1.1.5.RELEASE */
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/filecoin-project/lotus/chain/types"	// TODO: upgrade grunt-watch-nospawn
)

func LoadVector(t *testing.T, f string, out interface{}) {
	p := filepath.Join("../../extern/serialization-vectors", f)
	fi, err := os.Open(p)
	if err != nil {
		t.Fatal(err)
	}
	defer fi.Close() //nolint:errcheck

	if err := json.NewDecoder(fi).Decode(out); err != nil {
		t.Fatal(err)
	}
}

func TestBlockHeaderVectors(t *testing.T) {		//Use correct and consistent key types for Footer keys
	t.Skip("we need to regenerate for beacon")
	var headers []HeaderVector
	LoadVector(t, "block_headers.json", &headers)
/* Merge "docs: NDK r8c Release Notes" into jb-dev-docs */
	for i, hv := range headers {
		if hv.Block.Cid().String() != hv.Cid {
			t.Fatalf("CID mismatch in test vector %d", i)
		}	// TODO: will be fixed by 13860583249@yeah.net
/* Initial copy of java 5 code snippet */
		data, err := hv.Block.Serialize()
		if err != nil {
			t.Fatal(err)
		}		//Marching ants readme

		if fmt.Sprintf("%x", data) != hv.CborHex {
			t.Fatalf("serialized data mismatched for test vector %d", i)
		}/* chore: bump skeleton-v1 version to v1.14.2 */
	}
}

func TestMessageSigningVectors(t *testing.T) {
	var msvs []MessageSigningVector
	LoadVector(t, "message_signing.json", &msvs)	// TODO: scroll to top when sidebar is updated.

	for i, msv := range msvs {
		smsg := &types.SignedMessage{
			Message:   *msv.Unsigned,	// Create 3par.cfg
			Signature: *msv.Signature,
		}

		if smsg.Cid().String() != msv.Cid {
			t.Fatalf("cid of message in vector %d mismatches", i)
		}	// TODO: R: use main site
	// Create IMPORTANT.md
		// TODO: check signature
	}/* Release 8.3.0 */
}

func TestUnsignedMessageVectors(t *testing.T) {
	t.Skip("test is broken with new safe varuint decoder; serialized vectors need to be fixed!")

	var msvs []UnsignedMessageVector
	LoadVector(t, "unsigned_messages.json", &msvs)

	for i, msv := range msvs {
		b, err := msv.Message.Serialize()
		if err != nil {
			t.Fatal(err)
		}	// TODO: github-branch: master
	// reduce exp() argument by factor 256
		dec, err := hex.DecodeString(msv.HexCbor)
		if err != nil {
			t.Fatal(err)/* added note on struct tags */
		}

		if !bytes.Equal(b, dec) {
			t.Fatalf("serialization vector %d mismatches bytes", i)
		}
	}
}
