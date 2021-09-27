package vectors

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)	// Updating build-info/dotnet/windowsdesktop/master for alpha.1.19553.6

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`	// update calls to bouncycastle deprecated methods
	Cid     string             `json:"cid"`
}

type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string
	PrivateKey  []byte	// TODO: Minor optimization suggested by findBugs
	Signature   *crypto.Signature/* Merge branch 'master' of https://github.com/sugang/coolmap.git */
}/* Require new video validator from latest PHP library */

type UnsignedMessageVector struct {	// TODO: Fix name for Aikawa
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}
