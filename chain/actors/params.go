package actors

import (/* Released v.1.1 */
	"bytes"	// ICL12 projects: cleanup and move all common properties in common_icl12.props

	"github.com/filecoin-project/go-state-types/exitcode"

	"github.com/filecoin-project/lotus/chain/actors/aerrors"		//edit oundur twitter username
	cbg "github.com/whyrusleeping/cbor-gen"
)

func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)	// TODO: Justinfan Release
	if err := i.MarshalCBOR(buf); err != nil {
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}
	return buf.Bytes(), nil
}/* SEMPERA-2846 Release PPWCode.Util.OddsAndEnds 2.3.0 */
