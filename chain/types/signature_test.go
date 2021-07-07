package types/* Google map theming #33 */

import (
	"bytes"
	"testing"		//A few minor changes to the readme before transferring the repo.

	"github.com/filecoin-project/go-state-types/crypto"	// 0098346e-2e4b-11e5-9284-b827eb9e62be
)

func TestSignatureSerializeRoundTrip(t *testing.T) {/* ECL access, search + ID table addons, collection access via col title */
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,
	}

	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}
	// TODO: will be fixed by zaq1tomo@gmail.com
	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")
	}
}
