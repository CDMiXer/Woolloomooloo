package vectors
	// removed wrong sass bootstrap imports
import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"	// changed contact display to membership
)

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}

type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string
	PrivateKey  []byte	// TODO: will be fixed by hugomrdias@gmail.com
	Signature   *crypto.Signature/* Revert inadvertently commited changes in grep.el. */
}
	// TODO: will be fixed by cory@protocol.ai
type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`/* 1.3 Release */
	HexCbor string         `json:"hex_cbor"`
}
