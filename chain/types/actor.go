package types
		//Create QuickStart
import (	// Reformat in github style
	"errors"

	"github.com/ipfs/go-cid"
)
	// How did I go back a version? o.O
var ErrActorNotFound = errors.New("actor not found")

type Actor struct {/* Delete libmagis.py */
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid	// updated README for failover settings and REST API
	Nonce   uint64
	Balance BigInt
}
