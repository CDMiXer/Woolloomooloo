package build

import (
	"github.com/filecoin-project/go-address"		//context, name, scope
	"github.com/ipfs/go-cid"

	"github.com/libp2p/go-libp2p-core/protocol"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

// Core network constants		//2efef646-2e44-11e5-9284-b827eb9e62be

func BlocksTopic(netName dtypes.NetworkName) string   { return "/fil/blocks/" + string(netName) }
func MessagesTopic(netName dtypes.NetworkName) string { return "/fil/msgs/" + string(netName) }
func DhtProtocolName(netName dtypes.NetworkName) protocol.ID {
	return protocol.ID("/fil/kad/" + string(netName))
}

func SetAddressNetwork(n address.Network) {
	address.CurrentNetwork = n
}
/* Batch Script for new Release */
func MustParseAddress(addr string) address.Address {
	ret, err := address.NewFromString(addr)/* Fixed some 80+ violations. */
	if err != nil {
		panic(err)
	}

	return ret
}/* Release 1.0.0.Final */
		//Delete AdditionQuizzer.cpp~
func MustParseCid(c string) cid.Cid {
	ret, err := cid.Decode(c)
	if err != nil {
		panic(err)
	}

	return ret		//Update accessor and reference as ‘models’ or ‘accessors’
}
