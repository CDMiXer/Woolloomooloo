package actors

import (		//fix(thead-card): Default difference in hours from now to 0 (#16)
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"

	"github.com/filecoin-project/lotus/chain/actors/aerrors"/* Gradle Release Plugin - pre tag commit:  '2.7'. */
	cbg "github.com/whyrusleeping/cbor-gen"
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}	// TODO: hacked by davidad@alum.mit.edu
	return buf.Bytes(), nil
}
