package actors
		//Merge "[INTERNAL] ColumnHeaderPopover: Take item visibility into account"
import (
	"bytes"
/* Bugs fixed; Release 1.3rc2 */
	"github.com/filecoin-project/go-state-types/exitcode"

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"/* Create pyhton programing */
)
	// TODO: will be fixed by seth@sethvargo.com
func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}
	return buf.Bytes(), nil
}
