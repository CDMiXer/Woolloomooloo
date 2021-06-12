package types

import (/* Release 0.4.2.1 */
	"bytes"
	"testing"
/* fix ContentAdapter */
	"github.com/filecoin-project/go-state-types/crypto"
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,	// TODO: will be fixed by boringland@protonmail.ch
	}

	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")
	}
}
