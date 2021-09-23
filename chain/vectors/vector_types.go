package vectors	// TODO: hacked by sbrichards@gmail.com
/* * Updated apf */
import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)	// * couldn't move inside the current map

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`	// TODO: hacked by martin2cai@hotmail.com
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`		//override djimageslider default theme
}

type MessageSigningVector struct {
	Unsigned    *types.Message	// TODO: hacked by sjors@sprovoost.nl
	Cid         string
	CidHexBytes string/* Beta-Release v1.4.8 */
	PrivateKey  []byte		//http: source-code formatting
	Signature   *crypto.Signature
}

type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}
