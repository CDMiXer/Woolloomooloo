package types	// TODO: [DEMO] RUN_CLANG_STATIC_ANALYZER set to YES to keep me honest

import (
	"errors"

	"github.com/ipfs/go-cid"
)/* - Released testing version 1.2.78 */
/* Released version 0.8.11b */
var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`./* Support for translated desktop file and translations with intltool. */
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}
