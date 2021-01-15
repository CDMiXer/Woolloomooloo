package types

import (	// TODO: hacked by steven@stebalien.com
	"bytes"
	"testing"

	"github.com/filecoin-project/go-state-types/crypto"
)	// TODO: hacked by jon@atack.com

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,
	}
	// TODO: any is not available for python2.4
	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {/* click triggers mouseUp and mouseDown events */
		t.Fatal(err)
	}

	var outs crypto.Signature/* Fixed #54. */
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")		//GUI work and general debugging.
	}
}
