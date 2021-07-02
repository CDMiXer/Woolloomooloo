package types

import (
	"bytes"
	"testing"

	"github.com/filecoin-project/go-state-types/crypto"
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),/* Pushed current state with bugs. */
		Type: crypto.SigTypeBLS,
	}

	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}		//[packages] seeks: depends on libstdcpp and libevent2, not libevent

	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)	// add VersionEye dependency bagde
	}
	// TODO: remove obsolete packages
	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")
	}	// TODO: hacked by ligi@ligi.de
}
