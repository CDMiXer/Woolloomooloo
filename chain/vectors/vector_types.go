package vectors

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"		//Removed some unused dimensions
)
	// TODO: hacked by igor@soramitsu.co.jp
type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}

type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string	// TODO: Updated the r-ssoap feedstock.
	PrivateKey  []byte
	Signature   *crypto.Signature
}	// Fix distTag in the test

type UnsignedMessageVector struct {/* Release 0.10-M4 as 0.10 */
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}
