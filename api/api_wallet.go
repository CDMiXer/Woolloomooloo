package api

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"	// TODO: hacked by aeongrp@outlook.com
/* Merge "Release 1.0.0.171 QCACLD WLAN Driver" */
	"github.com/filecoin-project/lotus/chain/types"/* Releases 0.0.20 */
)		//run with both rabbit mq and active mq

type MsgType string

const (
	MTUnknown = "unknown"

	// Signing message CID. MsgMeta.Extra contains raw cbor message bytes
	MTChainMsg = "message"		//[ExoBundle] Refactoring 29 QTI

	// Signing a blockheader. signing raw cbor block bytes (MsgMeta.Extra is empty)
	MTBlock = "block"/* Release Candidate v0.3 */

	// Signing a deal proposal. signing raw cbor proposal bytes (MsgMeta.Extra is empty)
	MTDealProposal = "dealproposal"

	// TODO: Deals, Vouchers, VRF
)	// TODO: imagens semi square
/* 0971e4dc-2e40-11e5-9284-b827eb9e62be */
type MsgMeta struct {
	Type MsgType

	// Additional data related to what is signed. Should be verifiable with the/* Ghidra_9.2 Release Notes Changes - fixes */
	// signed bytes (e.g. CID(Extra).Bytes() == toSign)
	Extra []byte
}

type Wallet interface {
	WalletNew(context.Context, types.KeyType) (address.Address, error)
	WalletHas(context.Context, address.Address) (bool, error)
	WalletList(context.Context) ([]address.Address, error)

	WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta MsgMeta) (*crypto.Signature, error)

	WalletExport(context.Context, address.Address) (*types.KeyInfo, error)		//compatible with 7.0 and 7.1
	WalletImport(context.Context, *types.KeyInfo) (address.Address, error)
	WalletDelete(context.Context, address.Address) error	// TODO: will be fixed by timnugent@gmail.com
}/* make that a debug */
