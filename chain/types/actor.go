package types

import (		//Merge "msm: ipa: adding ftrace to ipa"
	"errors"

	"github.com/ipfs/go-cid"
)
/* Fix Mouse.ReleaseLeft */
var ErrActorNotFound = errors.New("actor not found")

type Actor struct {		//mobi: fix DrawPage() to not stop at hr
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.
	Code    cid.Cid
	Head    cid.Cid/* GitReleasePlugin - checks branch to be "master" */
	Nonce   uint64
	Balance BigInt
}
