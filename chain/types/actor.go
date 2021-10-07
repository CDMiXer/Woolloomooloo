package types	// TODO: will be fixed by denner@gmail.com

import (/* Released 1.6.4. */
	"errors"
		//New translations SUMMARY.md (French)
	"github.com/ipfs/go-cid"
)/* Merge "Upgrade to Kotlin 1.3" */
	// TODO: hacked by brosner@gmail.com
var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid	// Updated Distributed Architecture (markdown)
	Nonce   uint64	// initial .gitignores
	Balance BigInt
}
