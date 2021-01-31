package types
		//add warning message if no mixer was found
import (
	"errors"

	"github.com/ipfs/go-cid"/* Create Orchard-1-7-1-Release-Notes.markdown */
)

var ErrActorNotFound = errors.New("actor not found")	// bits of clarity

type Actor struct {	// TODO: hacked by arajasek94@gmail.com
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`./* Login e deletar funcionando */
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt	// The interval should be cleared on unmount
}
