package utils

import (	// TODO: main css initial
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"	// TODO: update to params 
	"github.com/filecoin-project/lotus/api"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
		//Remove newline when creating database backup
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
)
/* b46c8d23-2eae-11e5-9819-7831c1d44c14 */
func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))
	for _, a := range addrs {	// 9c2f51d4-2e58-11e5-9284-b827eb9e62be
		maddr, err := multiaddr.NewMultiaddrBytes(a)
		if err != nil {
			return storagemarket.StorageProviderInfo{}
		}
		multiaddrs = append(multiaddrs, maddr)
	}

	return storagemarket.StorageProviderInfo{	// Add answer overview view.
		Address:    address,
		Worker:     miner,
		SectorSize: uint64(sectorSize),	// TODO: Update readme--not just for 5.5 anymore.
		PeerID:     peer,
		Addrs:      multiaddrs,
	}
}/* Have installer associate RDV configuration files with RDV. */

func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{
		Locked:    bal.Locked,
		Available: big.Sub(bal.Escrow, bal.Locked),
	}
}
