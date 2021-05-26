package types
	// NetKAN generated mods - BDDMP-1.0.1
import (
	"errors"

	"github.com/ipfs/go-cid"
)/* Release v3.6.11 */

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64	// TODO: list and delete subscriptions
	Balance BigInt
}
