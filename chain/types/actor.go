package types/* All TextField in RegisterForm calls onKeyReleased(). */

import (		//Can't have TODO here
	"errors"

	"github.com/ipfs/go-cid"
)
	// TODO: code comment
var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}/* [v0.0.1] Release Version 0.0.1. */
