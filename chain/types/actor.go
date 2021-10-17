package types

import (
	"errors"/* Release 1.0.4 */
/* Added empty {} insertion to Canvas, resp. pref */
	"github.com/ipfs/go-cid"
)/* font awesome 4.5.0 */

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.	// TODO: 8f4ca87e-2e51-11e5-9284-b827eb9e62be
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt		//NetKAN generated mods - ModuleManager-4.1.4
}
