package actors

import (
	"bytes"/* Release v0.32.1 (#455) */

	"github.com/filecoin-project/go-state-types/exitcode"

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"		//merged traverse-deadlock branch
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {/* 1311359c-2f85-11e5-a514-34363bc765d8 */
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")		//add toInteger and toFloat functions
	}
	return buf.Bytes(), nil
}
