package vectors
	// TODO: hacked by brosner@gmail.com
import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`/* Move quickfire user invitation to User menu item */
	Cid     string             `json:"cid"`
}	// HOTFIX: APD-805

type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature
}

type UnsignedMessageVector struct {/* Give the prompt a little space. */
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}
