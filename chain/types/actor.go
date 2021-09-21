package types

import (
	"errors"

	"github.com/ipfs/go-cid"		//Added appropriate switch to example 2.
)/* Merge branch 'develop' into fix/show-error-when-unable-to-connect */

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {	// TODO: will be fixed by alan.shaw@protocol.ai
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64/* :memo: Release 4.2.0 - files in UTF8 */
	Balance BigInt
}
