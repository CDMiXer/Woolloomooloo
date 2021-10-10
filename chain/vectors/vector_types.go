package vectors/* Updated to version 1.2.0 */

import (
	"github.com/filecoin-project/go-state-types/crypto"/* detach while reading cache */
	"github.com/filecoin-project/lotus/chain/types"	// TODO: Merge "wlan: low throughput regression fix"
)	// TODO: Merge "Remove TODO comments in MCV" into androidx-master-dev

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`		//Update notification system
	Cid     string             `json:"cid"`
}

type MessageSigningVector struct {	// LI-USB programming working
	Unsigned    *types.Message
	Cid         string/* Release of version 0.1.1 */
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature
}/* also fixed saturation calc in color conversion  */

type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}
