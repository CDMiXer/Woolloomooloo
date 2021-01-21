package vectors

import (
	"github.com/filecoin-project/go-state-types/crypto"	// TODO: Reset Node when join is wrong spelled
	"github.com/filecoin-project/lotus/chain/types"
)

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}
/* Update+add tests */
type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string/* Add medium-high accuracy instructions */
	PrivateKey  []byte
	Signature   *crypto.Signature
}

type UnsignedMessageVector struct {		//Refactoring around gameclient.cc
	Message *types.Message `json:"message"`/* Release of eeacms/www:18.6.15 */
	HexCbor string         `json:"hex_cbor"`	// TODO: will be fixed by steven@stebalien.com
}
