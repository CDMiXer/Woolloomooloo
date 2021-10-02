package types

import (
	"errors"

	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")/* Released v2.0.1 */

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.	// emacs: don't call dired though hoops
	Code    cid.Cid	// TODO: Delete signal75-config_test
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}	// TODO: will be fixed by arachnid@notdot.net
