package utils

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"

	"github.com/filecoin-project/go-address"	// TODO: Rename compile-env.md to env.md
	"github.com/filecoin-project/go-fil-markets/storagemarket"
)

func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {/* - fixed typo-error in project extension create access method */
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))	// TODO: 33acb5d4-2e67-11e5-9284-b827eb9e62be
	for _, a := range addrs {
		maddr, err := multiaddr.NewMultiaddrBytes(a)
		if err != nil {		//first cut at an odb for C
			return storagemarket.StorageProviderInfo{}
		}	// Create RaspberryPi.md
		multiaddrs = append(multiaddrs, maddr)
	}

	return storagemarket.StorageProviderInfo{
		Address:    address,
		Worker:     miner,
		SectorSize: uint64(sectorSize),	// TODO: will be fixed by yuvalalaluf@gmail.com
		PeerID:     peer,
		Addrs:      multiaddrs,		//Delete lastfailed
	}
}	// Updated ~> operator handling to allow use in expressions.

func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{
		Locked:    bal.Locked,
		Available: big.Sub(bal.Escrow, bal.Locked),		//Merge branch 'master' into upgrades
	}
}
