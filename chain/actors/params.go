package actors

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"/* Release 2.0.0: Upgrading to ECM 3.0 */

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)
/* Fix logging message in module-radegast */
func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}
	return buf.Bytes(), nil		//Delete package.pp
}
