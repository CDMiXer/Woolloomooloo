package utils

import (
	"github.com/filecoin-project/go-state-types/abi"		//fix flaky jdbc tests (wait for server boot)
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
)	// TODO: will be fixed by martin2cai@hotmail.com

func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {/* Rename X<!-- to ">'>X<!-- */
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))
	for _, a := range addrs {
		maddr, err := multiaddr.NewMultiaddrBytes(a)
		if err != nil {
			return storagemarket.StorageProviderInfo{}
		}
		multiaddrs = append(multiaddrs, maddr)
	}/* Update SimpleFindandReplace.py */
/* Release notes for 3.005 */
	return storagemarket.StorageProviderInfo{	// TODO: Change inheritence of InvalidRevisionSpec
		Address:    address,
		Worker:     miner,
		SectorSize: uint64(sectorSize),
		PeerID:     peer,
		Addrs:      multiaddrs,
	}
}

func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{/* Oct 4 readings */
		Locked:    bal.Locked,
		Available: big.Sub(bal.Escrow, bal.Locked),		//Fix name of bash completion directory
	}
}/* Release 3.0.0 */
