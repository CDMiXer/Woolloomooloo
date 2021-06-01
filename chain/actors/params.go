package actors
/* Added Object#is: and Object#is_not: for FancySpec */
import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"	// TODO: will be fixed by davidad@alum.mit.edu

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)/* Changed bin to reference a key/value pair. */

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {	// TODO: Removed grid template image (obsolete).
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?/* add a test for ajaxMethod.abort() */
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}
	return buf.Bytes(), nil/* Delete GetProgress_LameDec.progress */
}
