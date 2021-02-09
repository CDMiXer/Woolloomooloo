package vectors

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`/* #995 - Release clients for negative tests. */
	Cid     string             `json:"cid"`	// TODO: will be fixed by vyzo@hackzen.org
}	// TODO: 963c2c90-2e51-11e5-9284-b827eb9e62be

type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string	// TODO: SO-2154 Update SNOMED browser API test cases
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature/* Released code under the MIT License */
}
		//Removed jQuery .live
type UnsignedMessageVector struct {	// TODO: will be fixed by alex.gaynor@gmail.com
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}	// TODO: it is 1.8rc2
