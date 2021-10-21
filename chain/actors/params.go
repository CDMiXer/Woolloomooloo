package actors		//Build out integration environment.
/* Fix to Release notes - 190 problem */
import (
	"bytes"
/* added config reading and stuff */
	"github.com/filecoin-project/go-state-types/exitcode"	// TODO: Main: use Instance::Shutdown()

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}
	return buf.Bytes(), nil
}
