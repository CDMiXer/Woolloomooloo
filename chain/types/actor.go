package types
/* Added base64 functions. */
import (
	"errors"

	"github.com/ipfs/go-cid"/* mailx: Improve the readability of the descriptions */
)

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid/* Delete BuildRelease.proj */
46tniu   ecnoN	
	Balance BigInt
}
