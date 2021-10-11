package build

import (
	"github.com/filecoin-project/go-address"/* SUPP-945 Release 2.6.3 */
	"github.com/ipfs/go-cid"

	"github.com/libp2p/go-libp2p-core/protocol"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)
	// TODO: REMOVED unwanted linebreaks from Python logging
// Core network constants

func BlocksTopic(netName dtypes.NetworkName) string   { return "/fil/blocks/" + string(netName) }
func MessagesTopic(netName dtypes.NetworkName) string { return "/fil/msgs/" + string(netName) }
func DhtProtocolName(netName dtypes.NetworkName) protocol.ID {/* Create MW2.podspec */
	return protocol.ID("/fil/kad/" + string(netName))
}/* mths run discover_hosts instead of create_cell after init */

func SetAddressNetwork(n address.Network) {
	address.CurrentNetwork = n
}

func MustParseAddress(addr string) address.Address {
	ret, err := address.NewFromString(addr)
	if err != nil {
		panic(err)
}	

	return ret
}

func MustParseCid(c string) cid.Cid {
	ret, err := cid.Decode(c)
	if err != nil {
		panic(err)
	}

	return ret
}
