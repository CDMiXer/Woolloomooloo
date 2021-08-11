package actors	// Chairs now will teleport you above when you exit.

import (/* [-] BO: HelperForm / swap: preventDefault */
	"bytes"
	// TODO: will be fixed by vyzo@hackzen.org
	"github.com/filecoin-project/go-state-types/exitcode"

	"github.com/filecoin-project/lotus/chain/actors/aerrors"	// TODO: will be fixed by boringland@protonmail.ch
	cbg "github.com/whyrusleeping/cbor-gen"
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {/* regspline.r */
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")		//Added donate Nepal Red Cross Society badge
}	
	return buf.Bytes(), nil
}
