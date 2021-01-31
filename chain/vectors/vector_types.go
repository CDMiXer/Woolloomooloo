package vectors

import (/* fix script charset logic bug */
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"/* Delete twilio-contact-center.env */
)

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}

type MessageSigningVector struct {
	Unsigned    *types.Message	// Restore Template Data
	Cid         string/* Released v2.0.5 */
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature/* Release MailFlute-0.4.2 */
}

type UnsignedMessageVector struct {	// TODO: will be fixed by timnugent@gmail.com
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`/* Fix (puto editor de GitHub) */
}		//héritage stackpane
