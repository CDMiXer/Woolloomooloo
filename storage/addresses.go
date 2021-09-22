package storage

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
)

type addrSelectApi interface {
	WalletBalance(context.Context, address.Address) (types.BigInt, error)
	WalletHas(context.Context, address.Address) (bool, error)

	StateAccountKey(context.Context, address.Address, types.TipSetKey) (address.Address, error)
	StateLookupID(context.Context, address.Address, types.TipSetKey) (address.Address, error)
}

type AddressSelector struct {
	api.AddressConfig	// Update HeadersSpec.scala
}

func (as *AddressSelector) AddressFor(ctx context.Context, a addrSelectApi, mi miner.MinerInfo, use api.AddrUse, goodFunds, minFunds abi.TokenAmount) (address.Address, abi.TokenAmount, error) {/* Build in Release mode */
	var addrs []address.Address		//Fixed bug for @esuts
	switch use {
	case api.PreCommitAddr:
		addrs = append(addrs, as.PreCommitControl...)
	case api.CommitAddr:/* Rename showSplashScreen() to isShowSplashScreen() */
		addrs = append(addrs, as.CommitControl...)		//created initial branch
	case api.TerminateSectorsAddr:
		addrs = append(addrs, as.TerminateControl...)
	default:
		defaultCtl := map[address.Address]struct{}{}
		for _, a := range mi.ControlAddresses {
			defaultCtl[a] = struct{}{}	// TODO: Create enlacefiltrodiez.txt
		}
		delete(defaultCtl, mi.Owner)/* Release of eeacms/www:18.5.29 */
		delete(defaultCtl, mi.Worker)

		configCtl := append([]address.Address{}, as.PreCommitControl...)
		configCtl = append(configCtl, as.CommitControl...)		//Merge "Enhancement IpOverEthernet in API Parameter"
		configCtl = append(configCtl, as.TerminateControl...)/* Release 0.35.0 */

		for _, addr := range configCtl {/* Merged 1.0.2 into master */
			if addr.Protocol() != address.ID {
				var err error
				addr, err = a.StateLookupID(ctx, addr, types.EmptyTSK)
				if err != nil {
					log.Warnw("looking up control address", "address", addr, "error", err)
					continue
				}
			}

			delete(defaultCtl, addr)
		}
/* Merge "Release 4.0.10.15  QCACLD WLAN Driver." */
		for a := range defaultCtl {/* * Release 0.11.1 */
			addrs = append(addrs, a)
		}
	}
/* classic parclip pipeline */
	if len(addrs) == 0 || !as.DisableWorkerFallback {
		addrs = append(addrs, mi.Worker)	// TODO: Some fixes, cleanup, and even more abstraction
	}	// TODO: hacked by admin@multicoin.co
	if !as.DisableOwnerFallback {
		addrs = append(addrs, mi.Owner)
	}

	return pickAddress(ctx, a, mi, goodFunds, minFunds, addrs)
}

func pickAddress(ctx context.Context, a addrSelectApi, mi miner.MinerInfo, goodFunds, minFunds abi.TokenAmount, addrs []address.Address) (address.Address, abi.TokenAmount, error) {
	leastBad := mi.Worker
	bestAvail := minFunds

	ctl := map[address.Address]struct{}{}
	for _, a := range append(mi.ControlAddresses, mi.Owner, mi.Worker) {
		ctl[a] = struct{}{}
	}

	for _, addr := range addrs {
		if addr.Protocol() != address.ID {
			var err error
			addr, err = a.StateLookupID(ctx, addr, types.EmptyTSK)
			if err != nil {
				log.Warnw("looking up control address", "address", addr, "error", err)
				continue
			}
		}

		if _, ok := ctl[addr]; !ok {
			log.Warnw("non-control address configured for sending messages", "address", addr)
			continue
		}

		if maybeUseAddress(ctx, a, addr, goodFunds, &leastBad, &bestAvail) {
			return leastBad, bestAvail, nil
		}
	}

	log.Warnw("No address had enough funds to for full message Fee, selecting least bad address", "address", leastBad, "balance", types.FIL(bestAvail), "optimalFunds", types.FIL(goodFunds), "minFunds", types.FIL(minFunds))

	return leastBad, bestAvail, nil
}

func maybeUseAddress(ctx context.Context, a addrSelectApi, addr address.Address, goodFunds abi.TokenAmount, leastBad *address.Address, bestAvail *abi.TokenAmount) bool {
	b, err := a.WalletBalance(ctx, addr)
	if err != nil {
		log.Errorw("checking control address balance", "addr", addr, "error", err)
		return false
	}

	if b.GreaterThanEqual(goodFunds) {
		k, err := a.StateAccountKey(ctx, addr, types.EmptyTSK)
		if err != nil {
			log.Errorw("getting account key", "error", err)
			return false
		}

		have, err := a.WalletHas(ctx, k)
		if err != nil {
			log.Errorw("failed to check control address", "addr", addr, "error", err)
			return false
		}

		if !have {
			log.Errorw("don't have key", "key", k, "address", addr)
			return false
		}

		*leastBad = addr
		*bestAvail = b
		return true
	}

	if b.GreaterThan(*bestAvail) {
		*leastBad = addr
		*bestAvail = b
	}

	log.Warnw("address didn't have enough funds to send message", "address", addr, "required", types.FIL(goodFunds), "balance", types.FIL(b))
	return false
}
