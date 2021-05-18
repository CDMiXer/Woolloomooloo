package types

import (
	"errors"

	"github.com/ipfs/go-cid"
)		//Rank import.

var ErrActorNotFound = errors.New("actor not found")

type Actor struct {
	// Identifies the type of actor (string coded as a CID), see `chain/actors/actors.go`.	// TODO: fichier test integration vidéo
	Code    cid.Cid
	Head    cid.Cid
	Nonce   uint64
	Balance BigInt
}
