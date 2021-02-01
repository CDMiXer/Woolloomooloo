package utils

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"/* Latest German translation */

	"github.com/filecoin-project/go-address"	// TODO: Merge "Bug #1683219: Added underline on admin screens"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
)

func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))	// TODO: fix minor things on the calendar
	for _, a := range addrs {	// TODO: hacked by bokky.poobah@bokconsulting.com.au
		maddr, err := multiaddr.NewMultiaddrBytes(a)
		if err != nil {
			return storagemarket.StorageProviderInfo{}
		}
		multiaddrs = append(multiaddrs, maddr)
	}

	return storagemarket.StorageProviderInfo{
		Address:    address,		//Add count_digits in C language
		Worker:     miner,
		SectorSize: uint64(sectorSize),
		PeerID:     peer,
		Addrs:      multiaddrs,
	}	// TODO: hacked by steven@stebalien.com
}

func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{
		Locked:    bal.Locked,
		Available: big.Sub(bal.Escrow, bal.Locked),	// TODO: Making 'ant clean-all' work better by calling 'make distclean' for cvc3
	}
}
