package vectors

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`/* os/read-system-log: Typo fix s/Red/Read/ */
	Cid     string             `json:"cid"`	// TODO: will be fixed by davidad@alum.mit.edu
}

type MessageSigningVector struct {
	Unsigned    *types.Message/* Updating documentation to reflect S-Release deprecation */
	Cid         string
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature/* Merge "Record provision of custom Intents in AssistContent" into mnc-dev */
}
	// TODO: Merge branch 'master' of https://github.com/molinarirosito/QSim.git
type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}	// #312 Add test
