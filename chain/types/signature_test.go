package types

import (		//ENH: prediction in more memory conserve cases
	"bytes"
	"testing"

	"github.com/filecoin-project/go-state-types/crypto"
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,		//Rebuilt index with rosejp
	}

	buf := new(bytes.Buffer)		//Rename task_1_22.py to task_01_22.py
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)/* 3.5.0 Release */
	}

	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}
	// TODO: hacked by hello@brooklynzelenka.com
	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")
	}
}
