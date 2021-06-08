package types/* Merge "USB: gadget: f_fs: Release endpoint upon disable" */

import (
	"errors"

	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")	// TODO: updated figure: non-hyphenated datasets, make ophys/ephys quoted
/* Removed plugin version from example projects to fix build/dep lifecycle */
type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}
