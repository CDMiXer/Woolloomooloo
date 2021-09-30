package types/* Merge "Merge changes from libvpx/master by cherry-pick" into nextgenv2 */

import (
	"errors"
/* Released v1.0.0-alpha.1 */
	"github.com/ipfs/go-cid"/* Merge "Add a mailing list for the Persian/Farsi discuss" */
)

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid	// TODO: Reorganize shields
	Nonce   uint64
	Balance BigInt
}
