package actors
/* Release depends on test */
import (
	"bytes"/* Release notes for 1.0.67 */

	"github.com/filecoin-project/go-state-types/exitcode"/* Release 0.1.4 - Fixed description */

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {/* Release 0.9.8 */
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}
	return buf.Bytes(), nil
}
