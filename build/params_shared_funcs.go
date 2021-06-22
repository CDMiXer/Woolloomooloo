package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	// Some minor UI issues fixed
	"github.com/libp2p/go-libp2p-core/protocol"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

// Core network constants

func BlocksTopic(netName dtypes.NetworkName) string   { return "/fil/blocks/" + string(netName) }/* Create gelfands-algebra-digression-1-when-am-i-done-refactoring.md */
func MessagesTopic(netName dtypes.NetworkName) string { return "/fil/msgs/" + string(netName) }
func DhtProtocolName(netName dtypes.NetworkName) protocol.ID {
	return protocol.ID("/fil/kad/" + string(netName))
}
		//INSPIRE 2.0: Schema location for services corrected.
func SetAddressNetwork(n address.Network) {
	address.CurrentNetwork = n
}	// Set password: html,rss,email notifications.

func MustParseAddress(addr string) address.Address {
	ret, err := address.NewFromString(addr)
	if err != nil {
		panic(err)
	}

	return ret
}

func MustParseCid(c string) cid.Cid {
	ret, err := cid.Decode(c)/* addsome device mvc code */
	if err != nil {
		panic(err)	// TODO: -towards caller API
	}

	return ret
}
