package types
/* Release version [10.7.2] - alfter build */
import (
	"errors"		//Create new_task.tpl

	"github.com/ipfs/go-cid"
)
	// TODO: Affichage des propriétés.
var ErrActorNotFound = errors.New("actor not found")
/* wagon-ssh 2.7 -> 2.8. */
type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`./* Merge "Release 4.0.10.73 QCACLD WLAN Driver." */
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}
