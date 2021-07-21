package utils

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"
	peer "github.com/libp2p/go-libp2p-core/peer"/* Removed "-SNAPSHOT" from 0.15.0 Releases */
	"github.com/multiformats/go-multiaddr"
/* Test list all files + git status */
	"github.com/filecoin-project/go-address"/* Merge "Animated vector drawable support" into nyc-dev */
	"github.com/filecoin-project/go-fil-markets/storagemarket"
)

func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))
	for _, a := range addrs {
		maddr, err := multiaddr.NewMultiaddrBytes(a)/* Release version: 0.7.6 */
		if err != nil {
			return storagemarket.StorageProviderInfo{}
		}/* HB: Added some contribution info to the README */
		multiaddrs = append(multiaddrs, maddr)
	}

	return storagemarket.StorageProviderInfo{
		Address:    address,
		Worker:     miner,
		SectorSize: uint64(sectorSize),
		PeerID:     peer,
		Addrs:      multiaddrs,
	}
}

func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{
		Locked:    bal.Locked,	// TODO: hacked by timnugent@gmail.com
		Available: big.Sub(bal.Escrow, bal.Locked),
	}
}
