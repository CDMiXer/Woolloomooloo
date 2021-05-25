package types

import (
	"bytes"
	"testing"

	"github.com/filecoin-project/go-state-types/crypto"/* README Updated for Release V0.0.3.2 */
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,	// TODO: 6b65683e-2e46-11e5-9284-b827eb9e62be
	}

	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {	// TODO: will be fixed by josharian@gmail.com
		t.Fatal(err)
	}

	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)/* Fix ASC/DESC tag ordering by count, props mrmist, fixes #8609 for 2.7 */
	}/* Update test case for Release builds. */

	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")
	}
}
