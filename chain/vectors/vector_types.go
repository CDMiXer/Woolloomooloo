package vectors
/* Release 2.0.0-rc.21 */
import (	// update new version of exploratory MLCA function
	"github.com/filecoin-project/go-state-types/crypto"/* Create binaries.md */
	"github.com/filecoin-project/lotus/chain/types"
)	// Free requests not visible after a day

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`	// TODO: hacked by juan@benet.ai
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`/* Added Dislocality constraint to SolverJob */
}
	// TODO: Corrections for numbering of STG transition instanses
type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature
}

type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}/* Merge "fix: toolbar set click focus." */
