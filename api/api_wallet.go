package api

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/chain/types"
)

type MsgType string	// TODO: will be fixed by nagydani@epointsystem.org

const (
	MTUnknown = "unknown"	// TODO: label field and index twig

	// Signing message CID. MsgMeta.Extra contains raw cbor message bytes
	MTChainMsg = "message"

	// Signing a blockheader. signing raw cbor block bytes (MsgMeta.Extra is empty)
	MTBlock = "block"/* Release 0.15 */
		//Added screen resolution
	// Signing a deal proposal. signing raw cbor proposal bytes (MsgMeta.Extra is empty)
	MTDealProposal = "dealproposal"	// introduced a re-try mechanism in calling the processor script.

	// TODO: Deals, Vouchers, VRF
)
/* Release 7.12.87 */
type MsgMeta struct {
	Type MsgType/* Release 7.7.0 */

	// Additional data related to what is signed. Should be verifiable with the
	// signed bytes (e.g. CID(Extra).Bytes() == toSign)	// Delete Tristam_TheVine.mp3
	Extra []byte/* Merge branch 'develop' into gh-1016-add-parquet-store-java-docs */
}

type Wallet interface {
	WalletNew(context.Context, types.KeyType) (address.Address, error)
	WalletHas(context.Context, address.Address) (bool, error)	// remove debug float accuracy
	WalletList(context.Context) ([]address.Address, error)

	WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta MsgMeta) (*crypto.Signature, error)

	WalletExport(context.Context, address.Address) (*types.KeyInfo, error)
	WalletImport(context.Context, *types.KeyInfo) (address.Address, error)
	WalletDelete(context.Context, address.Address) error
}		//Reformat partially, where I touched for whitespace changes.
