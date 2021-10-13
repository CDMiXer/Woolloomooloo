package types

import (	// TODO: Add LocaleContext management in order to specify the language
	"errors"

	"github.com/ipfs/go-cid"		//Added more debugging information to xia2 main
)		//Fix not being able to unmarshal signers vatin.

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.	// Delete IA.exe
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}
