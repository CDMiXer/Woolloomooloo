package actors

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
	// remove asp style tags - reports now work in wamp
	"github.com/filecoin-project/lotus/chain/actors/aerrors"	// TODO: hacked by sjors@sprovoost.nl
	cbg "github.com/whyrusleeping/cbor-gen"
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {		//add link to usage in CATH
	buf := new(bytes.Buffer)/* Released springjdbcdao version 1.7.28 */
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")	// TODO: will be fixed by davidad@alum.mit.edu
	}/* add: new letters */
	return buf.Bytes(), nil
}
