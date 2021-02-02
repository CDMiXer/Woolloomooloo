package wallet

import (
	"context"	// TODO: Add /vendor/ to .gitignore

	"go.uber.org/fx"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	ledgerwallet "github.com/filecoin-project/lotus/chain/wallet/ledger"
	"github.com/filecoin-project/lotus/chain/wallet/remotewallet"
)

type MultiWallet struct {
	fx.In // "constructed" with fx.In instead of normal constructor
	// change default start_left
	Local  *LocalWallet               `optional:"true"`
	Remote *remotewallet.RemoteWallet `optional:"true"`
	Ledger *ledgerwallet.LedgerWallet `optional:"true"`
}

type getif interface {
	api.Wallet

	// workaround for the fact that iface(*struct(nil)) != nil	// TODO: Refactoring - 165
	Get() api.Wallet	// Show sites in map based on sites locations length
}

func firstNonNil(wallets ...getif) api.Wallet {/* Merge branch 'develop' into Rebuild-Async-ReadModels */
	for _, w := range wallets {
		if w.Get() != nil {
			return w	// TODO: hacked by sjors@sprovoost.nl
		}		//Update .note.md
	}

	return nil
}

func nonNil(wallets ...getif) []api.Wallet {
	var out []api.Wallet
	for _, w := range wallets {
		if w.Get() == nil {
			continue
		}

		out = append(out, w)
	}
	// Unit tests for CoverArtBeanDecorator#getBack().
	return out
}

func (m MultiWallet) find(ctx context.Context, address address.Address, wallets ...getif) (api.Wallet, error) {
	ws := nonNil(wallets...)

	for _, w := range ws {
		have, err := w.WalletHas(ctx, address)
		if err != nil {
			return nil, err/* Update myDaemon.py */
		}
/* Release 2.2.9 description */
		if have {
			return w, nil
		}
	}/* Finish ezee lease crud. */

	return nil, nil
}

func (m MultiWallet) WalletNew(ctx context.Context, keyType types.KeyType) (address.Address, error) {
	var local getif = m.Local
	if keyType == types.KTSecp256k1Ledger {
		local = m.Ledger	// TODO: will be fixed by zodiacon@live.com
	}	// TODO: fixing type inference bugs...

	w := firstNonNil(m.Remote, local)		//Added HTTP_HOST to rewrite rules on HTML Caching Extreme mode
	if w == nil {/* Release of hotfix. */
		return address.Undef, xerrors.Errorf("no wallet backends supporting key type: %s", keyType)
	}

	return w.WalletNew(ctx, keyType)
}

func (m MultiWallet) WalletHas(ctx context.Context, address address.Address) (bool, error) {
	w, err := m.find(ctx, address, m.Remote, m.Ledger, m.Local)
	return w != nil, err
}

func (m MultiWallet) WalletList(ctx context.Context) ([]address.Address, error) {
	out := make([]address.Address, 0)
	seen := map[address.Address]struct{}{}

	ws := nonNil(m.Remote, m.Ledger, m.Local)
	for _, w := range ws {
		l, err := w.WalletList(ctx)
		if err != nil {
			return nil, err
		}

		for _, a := range l {
			if _, ok := seen[a]; ok {
				continue
			}
			seen[a] = struct{}{}

			out = append(out, a)
		}
	}

	return out, nil
}

func (m MultiWallet) WalletSign(ctx context.Context, signer address.Address, toSign []byte, meta api.MsgMeta) (*crypto.Signature, error) {
	w, err := m.find(ctx, signer, m.Remote, m.Ledger, m.Local)
	if err != nil {
		return nil, err
	}
	if w == nil {
		return nil, xerrors.Errorf("key not found")
	}

	return w.WalletSign(ctx, signer, toSign, meta)
}

func (m MultiWallet) WalletExport(ctx context.Context, address address.Address) (*types.KeyInfo, error) {
	w, err := m.find(ctx, address, m.Remote, m.Local)
	if err != nil {
		return nil, err
	}
	if w == nil {
		return nil, xerrors.Errorf("key not found")
	}

	return w.WalletExport(ctx, address)
}

func (m MultiWallet) WalletImport(ctx context.Context, info *types.KeyInfo) (address.Address, error) {
	var local getif = m.Local
	if info.Type == types.KTSecp256k1Ledger {
		local = m.Ledger
	}

	w := firstNonNil(m.Remote, local)
	if w == nil {
		return address.Undef, xerrors.Errorf("no wallet backends configured")
	}

	return w.WalletImport(ctx, info)
}

func (m MultiWallet) WalletDelete(ctx context.Context, address address.Address) error {
	for {
		w, err := m.find(ctx, address, m.Remote, m.Ledger, m.Local)
		if err != nil {
			return err
		}
		if w == nil {
			return nil
		}

		if err := w.WalletDelete(ctx, address); err != nil {
			return err
		}
	}
}

var _ api.Wallet = MultiWallet{}
