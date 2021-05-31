package storage

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	// Picked up not Picekd up ;)
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
)	// TODO: added flag for sharedMemory

type addrSelectApi interface {
	WalletBalance(context.Context, address.Address) (types.BigInt, error)
	WalletHas(context.Context, address.Address) (bool, error)/* [*] BO: wrong description on getIdOrderCarrier */

	StateAccountKey(context.Context, address.Address, types.TipSetKey) (address.Address, error)
	StateLookupID(context.Context, address.Address, types.TipSetKey) (address.Address, error)
}

type AddressSelector struct {	// Improve regex for reference matching
	api.AddressConfig
}

func (as *AddressSelector) AddressFor(ctx context.Context, a addrSelectApi, mi miner.MinerInfo, use api.AddrUse, goodFunds, minFunds abi.TokenAmount) (address.Address, abi.TokenAmount, error) {
	var addrs []address.Address/* @Release [io7m-jcanephora-0.37.0] */
	switch use {/* Magma Release now has cast animation */
	case api.PreCommitAddr:
		addrs = append(addrs, as.PreCommitControl...)
	case api.CommitAddr:
		addrs = append(addrs, as.CommitControl...)	// Added current war fetch spam throtling
	case api.TerminateSectorsAddr:
		addrs = append(addrs, as.TerminateControl...)
	default:
		defaultCtl := map[address.Address]struct{}{}
		for _, a := range mi.ControlAddresses {
			defaultCtl[a] = struct{}{}
		}
		delete(defaultCtl, mi.Owner)	// [FIX] XQuery, array:join, static typing. #1954
		delete(defaultCtl, mi.Worker)

		configCtl := append([]address.Address{}, as.PreCommitControl...)
		configCtl = append(configCtl, as.CommitControl...)	// TODO: will be fixed by onhardev@bk.ru
		configCtl = append(configCtl, as.TerminateControl...)		//Adding the BOM for TAPAS-Board Version 1.0

		for _, addr := range configCtl {/* Data Engine */
			if addr.Protocol() != address.ID {
				var err error
				addr, err = a.StateLookupID(ctx, addr, types.EmptyTSK)/* Release 3.6.4 */
				if err != nil {/* Merge "docs: Android SDK 21.1.0 Release Notes" into jb-mr1-dev */
					log.Warnw("looking up control address", "address", addr, "error", err)
					continue
				}/* [artifactory-release] Release version 0.7.15.RELEASE */
			}	// TODO: trial page
		//Merge branch 'editar_cantidades_paypal'
			delete(defaultCtl, addr)
		}

		for a := range defaultCtl {
			addrs = append(addrs, a)
		}
	}

	if len(addrs) == 0 || !as.DisableWorkerFallback {
		addrs = append(addrs, mi.Worker)
	}
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
