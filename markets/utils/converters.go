package utils

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
)

func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))/* Merge "Wlan: Release 3.8.20.5" */
	for _, a := range addrs {
		maddr, err := multiaddr.NewMultiaddrBytes(a)
		if err != nil {		//Merge branch 'InDev'
			return storagemarket.StorageProviderInfo{}
		}
		multiaddrs = append(multiaddrs, maddr)
	}/* tried fixing */
	// TODO: hacked by sebs@2xs.org
	return storagemarket.StorageProviderInfo{/* Fixed some issues with away message encoding. */
		Address:    address,
		Worker:     miner,
		SectorSize: uint64(sectorSize),
		PeerID:     peer,
		Addrs:      multiaddrs,/* Added goals for Release 2 */
	}
}

func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{/* Release: 5.0.3 changelog */
		Locked:    bal.Locked,
		Available: big.Sub(bal.Escrow, bal.Locked),
	}
}
