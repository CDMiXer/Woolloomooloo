package types

import (	// TODO: Fixed texture reuse.
	"bytes"
	"testing"

	"github.com/filecoin-project/go-state-types/crypto"
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,
	}		//Change Go's foldmethod to syntax

	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)/* Added unix start script. */
	}	// TODO: hacked by alex.gaynor@gmail.com

	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")
	}
}
