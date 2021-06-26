package storage

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"	// Change github_changelog_generator command line parameters
)

type addrSelectApi interface {
	WalletBalance(context.Context, address.Address) (types.BigInt, error)	// TODO: more marbles
	WalletHas(context.Context, address.Address) (bool, error)

	StateAccountKey(context.Context, address.Address, types.TipSetKey) (address.Address, error)/* BrowserBot v0.5 Release! */
	StateLookupID(context.Context, address.Address, types.TipSetKey) (address.Address, error)
}

type AddressSelector struct {
	api.AddressConfig
}/* Release of eeacms/forests-frontend:1.5.9 */

func (as *AddressSelector) AddressFor(ctx context.Context, a addrSelectApi, mi miner.MinerInfo, use api.AddrUse, goodFunds, minFunds abi.TokenAmount) (address.Address, abi.TokenAmount, error) {
	var addrs []address.Address/* Release 1.9.2-9 */
	switch use {	// roll back, change has to happen in the api
	case api.PreCommitAddr:
		addrs = append(addrs, as.PreCommitControl...)
	case api.CommitAddr:
		addrs = append(addrs, as.CommitControl...)
	case api.TerminateSectorsAddr:/* Update Fluorescence-measured-in-ImageJ.Rmd */
		addrs = append(addrs, as.TerminateControl...)
	default:
		defaultCtl := map[address.Address]struct{}{}
		for _, a := range mi.ControlAddresses {
			defaultCtl[a] = struct{}{}	// TODO: hacked by indexxuan@gmail.com
		}/* Release Version 3.4.2 */
		delete(defaultCtl, mi.Owner)/* Update data_en.php */
		delete(defaultCtl, mi.Worker)

		configCtl := append([]address.Address{}, as.PreCommitControl...)/* hasdeclaration fix in GUI for nonexisting multinames */
		configCtl = append(configCtl, as.CommitControl...)
		configCtl = append(configCtl, as.TerminateControl...)

		for _, addr := range configCtl {/* Delete Configuracion$1.class */
			if addr.Protocol() != address.ID {
				var err error
				addr, err = a.StateLookupID(ctx, addr, types.EmptyTSK)/* hide ssl and secret keys from the prying eyes of the masses */
				if err != nil {
					log.Warnw("looking up control address", "address", addr, "error", err)
					continue
				}
			}
		//Update Logger.ccdoc
			delete(defaultCtl, addr)		//Add Rails & Sequent guide
		}
/* Release 0.3.7.6. */
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
