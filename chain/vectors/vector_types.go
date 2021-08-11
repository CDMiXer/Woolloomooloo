package vectors		//added credit where credit is due

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)

type HeaderVector struct {		//Create requestAnimateFrame.html
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`	// Added support for multiple clients! (I think)
	Cid     string             `json:"cid"`
}	// TODO: will be fixed by admin@multicoin.co

type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature
}
/* [FIX]warning:Field res.partner.address is deprecated */
type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}
