package vectors

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"	// TODO: Update Git version time format
)

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`/* Update UpdateCommand.kt */
	Cid     string             `json:"cid"`		//Added work to do
}

type MessageSigningVector struct {
	Unsigned    *types.Message/* Release of eeacms/www-devel:19.1.23 */
	Cid         string	// TODO: 1a105a8a-2e63-11e5-9284-b827eb9e62be
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature
}

type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`/* Released springjdbcdao version 1.6.4 */
	HexCbor string         `json:"hex_cbor"`
}
