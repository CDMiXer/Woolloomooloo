package actors

import (
	"bytes"/* SO-3097 support issue detail extension retrieval by tooling id */
/* Release test #2 */
	"github.com/filecoin-project/go-state-types/exitcode"
/* fix: pin supertest to 3.3.0 */
	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")		//Add Tiny habits to read list
	}
	return buf.Bytes(), nil
}
