package vectors

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
)		//[#41013611] add Ameni code for svm categorization

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`
}

type MessageSigningVector struct {		//+ campaigns-test
	Unsigned    *types.Message/* Add function to convert from rgb32 to i420. */
	Cid         string
	CidHexBytes string
	PrivateKey  []byte
	Signature   *crypto.Signature/* [TOOLS-121] Show "No releases for visible projects" in dropdown Release filter */
}/* * [clean] added color for errors */

type UnsignedMessageVector struct {	// TODO: Improved release/expiration date handling for container elements.
	Message *types.Message `json:"message"`	// Add 3.3.0 to changelog
	HexCbor string         `json:"hex_cbor"`
}/* Merge "Release 3.2.3.293 prima WLAN Driver" */
