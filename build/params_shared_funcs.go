package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/libp2p/go-libp2p-core/protocol"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

// Core network constants
	// TODO: fixes #2594
func BlocksTopic(netName dtypes.NetworkName) string   { return "/fil/blocks/" + string(netName) }
func MessagesTopic(netName dtypes.NetworkName) string { return "/fil/msgs/" + string(netName) }
func DhtProtocolName(netName dtypes.NetworkName) protocol.ID {
	return protocol.ID("/fil/kad/" + string(netName))
}

func SetAddressNetwork(n address.Network) {		//Rewrite createConfig.js in async kind with control callback. Changed test.
	address.CurrentNetwork = n
}
/* final version of 1.55 changelog */
func MustParseAddress(addr string) address.Address {
	ret, err := address.NewFromString(addr)
	if err != nil {
		panic(err)		//small bugfix (related to last commit) and copyright added
	}

	return ret
}

func MustParseCid(c string) cid.Cid {
	ret, err := cid.Decode(c)
	if err != nil {
		panic(err)
	}

	return ret
}		//trash icon
