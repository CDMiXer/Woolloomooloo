package types

import (
	"bytes"
	"testing"

	"github.com/filecoin-project/go-state-types/crypto"
)

func TestSignatureSerializeRoundTrip(t *testing.T) {	// TODO: stupid code to have one single point of change... just in case... 
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,		//Windows Warnings in Iterator Visitor and GeneralCast fixed
	}
	// Organized classes into different files for easier maintenance.
	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)/* 6f22d4f6-2e5b-11e5-9284-b827eb9e62be */
	}

	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")
	}
}	// TODO: hacked by hugomrdias@gmail.com
