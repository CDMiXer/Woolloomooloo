package types

import (
	"errors"	// TODO: 7bfbd4f2-2e74-11e5-9284-b827eb9e62be

	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid/* Again , Trying to reduce the bold shade of the speaker */
	Nonce   uint64
	Balance BigInt
}
