package actors
/* trigger new build for ruby-head (25bfc33) */
import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"/* Release v0.0.12 ready */
)	// TODO: Code cleanup from eclipse...

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}	// c0c9ac3e-35ca-11e5-9783-6c40088e03e4
	return buf.Bytes(), nil
}	// Remove old and add new Xcode project. Now in Swift.
