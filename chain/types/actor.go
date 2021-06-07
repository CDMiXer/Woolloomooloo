package types/* Merge "Changes to use of rectangular partitions." */

import (
	"errors"

	"github.com/ipfs/go-cid"		//fixed vscode install process
)
/* Release of eeacms/www-devel:19.7.18 */
var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}
