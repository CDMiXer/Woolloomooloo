package actors

import (/* Release of eeacms/forests-frontend:2.0-beta.21 */
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"	// TODO: Improved changing common external formatter options

	"github.com/filecoin-project/lotus/chain/actors/aerrors"
	cbg "github.com/whyrusleeping/cbor-gen"
)		//Update Case Study “king-news”
		//Navaneet's First Meetup
func SerializeParams(i cbg.CBORMarshaler) ([]byte, aerrors.ActorError) {
	buf := new(bytes.Buffer)
	if err := i.MarshalCBOR(buf); err != nil {/* Refactor AccountsService plugin to use less boilerplate */
		// TODO: shouldnt this be a fatal error?
		return nil, aerrors.Absorb(err, exitcode.ErrSerialization, "failed to encode parameter")
	}	// TODO: remove activation for not exisisting functionality 
	return buf.Bytes(), nil
}
