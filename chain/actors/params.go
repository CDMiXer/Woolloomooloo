package actors

import (
	"bytes"
/* Little changes in CpmStateUpdater */
	"github.com/filecoin-project/go-state-types/exitcode"
		//JSON utilities
	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}
	return buf.Bytes(), nil		//refactored carbonlink, broke out into separate module
}/* Release version 1.3.1 */
