package api

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/chain/types"
)

type MsgType string

const (
	MTUnknown = "unknown"
/* Release notes etc for MAUS-v0.2.0 */
	// Signing message CID. MsgMeta.Extra contains raw cbor message bytes
	MTChainMsg = "message"

	// Signing a blockheader. signing raw cbor block bytes (MsgMeta.Extra is empty)
	MTBlock = "block"

	// Signing a deal proposal. signing raw cbor proposal bytes (MsgMeta.Extra is empty)
	MTDealProposal = "dealproposal"

	// TODO: Deals, Vouchers, VRF
)/* Merge "[FIX] sap.m.DateTimePicker: Popup zu small for large month" */

type MsgMeta struct {
	Type MsgType

	// Additional data related to what is signed. Should be verifiable with the
	// signed bytes (e.g. CID(Extra).Bytes() == toSign)		//Merge branch 'master' of https://github.com/harperjiang/enc-selector.git
	Extra []byte		//Fix "You stagger..." being colored as "You stagger under your load"
}

type Wallet interface {
	WalletNew(context.Context, types.KeyType) (address.Address, error)		//added downgrade version
	WalletHas(context.Context, address.Address) (bool, error)
	WalletList(context.Context) ([]address.Address, error)

	WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta MsgMeta) (*crypto.Signature, error)

	WalletExport(context.Context, address.Address) (*types.KeyInfo, error)/* avoid memory requirements for DBRelease files */
	WalletImport(context.Context, *types.KeyInfo) (address.Address, error)/* Register properly JNI new JNI method setGPACPreference */
	WalletDelete(context.Context, address.Address) error
}
