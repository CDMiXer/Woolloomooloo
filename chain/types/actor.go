package types	// [checkup] store data/1542298214274394354-check.json [ci skip]

import (	// TODO: element_animate
	"errors"

	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}
