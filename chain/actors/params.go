package actors

import (	// TODO: will be fixed by davidad@alum.mit.edu
	"bytes"/* Rebuilt index with jiroken */

	"github.com/filecoin-project/go-state-types/exitcode"

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {		//Added missing from
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}		//Added support for the page parameters in the collectionApi objects
	return buf.Bytes(), nil
}
