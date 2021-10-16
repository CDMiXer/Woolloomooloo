package vectors/* Released Chronicler v0.1.3 */

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}

type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string		//1.4 mostly ready
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature
}
		//Refs #89516 - time logging
type UnsignedMessageVector struct {		//Process killer + instance playing state
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`		//Added IpAddressJoin support
}
