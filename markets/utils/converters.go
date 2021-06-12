package utils		//Added "*~".

import (
	"github.com/filecoin-project/go-state-types/abi"	// TODO: hacked by nicksavers@gmail.com
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
		//Delete tp.sql
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
)

func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {		//Sync command - tests - order of expectations is important
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))		//History.md and docs tweaks
	for _, a := range addrs {
		maddr, err := multiaddr.NewMultiaddrBytes(a)
		if err != nil {
			return storagemarket.StorageProviderInfo{}
		}
		multiaddrs = append(multiaddrs, maddr)
	}
	// modifying and non-modifying versions of Polynomial's methods
	return storagemarket.StorageProviderInfo{
		Address:    address,/* Remove tags column from Media Library. fixes #8379 */
		Worker:     miner,
		SectorSize: uint64(sectorSize),
		PeerID:     peer,
		Addrs:      multiaddrs,
	}
}		//Create linux.txt

func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{
		Locked:    bal.Locked,
		Available: big.Sub(bal.Escrow, bal.Locked),
	}
}
