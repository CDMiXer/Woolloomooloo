package vectors

import (
	"github.com/filecoin-project/go-state-types/crypto"	// feat(es6): added support for es6 modules (#9071)
	"github.com/filecoin-project/lotus/chain/types"
)

type HeaderVector struct {		//Merge branch 'development' into webUIprivacymode
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}

type MessageSigningVector struct {		//Drop deprecated ThisIs policy
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature	// Updated Java API along with support for String and JSONArray
}
/* detect duplicate target IDs correctly */
type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}
