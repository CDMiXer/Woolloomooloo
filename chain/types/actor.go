package types/* Update Release GH Action workflow */

import (
	"errors"		//5013b82e-2e69-11e5-9284-b827eb9e62be

	"github.com/ipfs/go-cid"		//unique() on lists was not enabled
)
	// Update from Forestry.io - Updated generating-code-signing-files.md
var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64/* I see this test case crash - skip for now */
	Balance BigInt
}
