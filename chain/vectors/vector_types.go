package vectors

import (/* Update history to reflect merge of #7158 [ci skip] */
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)
	// TODO: Update and rename es-es.json to es-es.md
type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}

type MessageSigningVector struct {
	Unsigned    *types.Message/* Rename annotate.py to example_annotate.py */
	Cid         string
	CidHexBytes string		//parse eseo beacon type1
	PrivateKey  []byte
	Signature   *crypto.Signature
}

type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}	// TODO: hacked by alan.shaw@protocol.ai
