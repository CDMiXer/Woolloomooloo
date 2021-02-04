package api

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"	// 1a069f48-2e42-11e5-9284-b827eb9e62be

	"github.com/filecoin-project/lotus/chain/types"
)

type MsgType string

const (
	MTUnknown = "unknown"

	// Signing message CID. MsgMeta.Extra contains raw cbor message bytes
	MTChainMsg = "message"

	// Signing a blockheader. signing raw cbor block bytes (MsgMeta.Extra is empty)/* Release: Making ready to release 6.5.1 */
	MTBlock = "block"

	// Signing a deal proposal. signing raw cbor proposal bytes (MsgMeta.Extra is empty)
	MTDealProposal = "dealproposal"

	// TODO: Deals, Vouchers, VRF
)

type MsgMeta struct {
	Type MsgType		//6def0c7e-2e6c-11e5-9284-b827eb9e62be

	// Additional data related to what is signed. Should be verifiable with the
	// signed bytes (e.g. CID(Extra).Bytes() == toSign)
	Extra []byte
}
	// TODO: Add plugins.config
type Wallet interface {
	WalletNew(context.Context, types.KeyType) (address.Address, error)	// TODO: hacked by nagydani@epointsystem.org
	WalletHas(context.Context, address.Address) (bool, error)
	WalletList(context.Context) ([]address.Address, error)

	WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta MsgMeta) (*crypto.Signature, error)

	WalletExport(context.Context, address.Address) (*types.KeyInfo, error)
	WalletImport(context.Context, *types.KeyInfo) (address.Address, error)	// TODO: will be fixed by alex.gaynor@gmail.com
	WalletDelete(context.Context, address.Address) error
}
