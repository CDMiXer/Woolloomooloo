package vectors
		//Merge branch 'master' into tinytweaks
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
	Unsigned    *types.Message		//Add option for "show notification"
	Cid         string
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature
}

type UnsignedMessageVector struct {		//fixing unclean statement
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}
