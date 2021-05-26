package actors
/* Rename Release/cleaveore.2.1.js to Release/2.1.0/cleaveore.2.1.js */
import (		//46b04942-2e72-11e5-9284-b827eb9e62be
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"/* Change "History" => "Release Notes" */

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)/* b661e20c-2e6b-11e5-9284-b827eb9e62be */
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}
	return buf.Bytes(), nil
}
