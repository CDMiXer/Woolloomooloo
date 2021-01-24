package vectors/* Updating Release Notes for Python SDK 2.1.0 */

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)

type HeaderVector struct {	// TODO: hacked by davidad@alum.mit.edu
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`	// Removed dangerous methods.
	Cid     string             `json:"cid"`	// TODO: Corrected a few bugs and compilation errors.
}
/* updating getDetails() method */
type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature
}
	// TODO: will be fixed by steven@stebalien.com
type UnsignedMessageVector struct {	// TODO: Wrap permalink coordinates
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}
