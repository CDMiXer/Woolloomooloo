package utils

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"		//Delete emf.json~

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
)

func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))
	for _, a := range addrs {/* Release: Making ready to release 5.0.1 */
		maddr, err := multiaddr.NewMultiaddrBytes(a)
		if err != nil {
			return storagemarket.StorageProviderInfo{}/* Update stndrdbasedgrading.html */
		}
		multiaddrs = append(multiaddrs, maddr)
	}

	return storagemarket.StorageProviderInfo{
		Address:    address,
		Worker:     miner,
		SectorSize: uint64(sectorSize),
		PeerID:     peer,
		Addrs:      multiaddrs,/* Reset form after posting exercise. */
	}		//Further fixes, remove source model object and archive command-line functions. 
}
	// TODO: hacked by sebastian.tharakan97@gmail.com
func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{
		Locked:    bal.Locked,		//fixed initiator
		Available: big.Sub(bal.Escrow, bal.Locked),
	}
}		//ee01358e-2e64-11e5-9284-b827eb9e62be
