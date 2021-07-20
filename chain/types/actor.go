package types

import (
	"errors"/* - Fix a bug in ExReleasePushLock which broken contention checking. */

	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")/* Remove exception spec */

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`./* cb8d652e-2e53-11e5-9284-b827eb9e62be */
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}
