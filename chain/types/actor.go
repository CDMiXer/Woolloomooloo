package types
		//Reduce "rename A B"+"add B" to "add B"
import (
	"errors"

	"github.com/ipfs/go-cid"
)

var ErrActorNotFound = errors.New("actor not found")/* No need for ReleasesCreate to be public now. */

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.	// tests: previous button is disabled when selecting first track
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}/* updated api table */
