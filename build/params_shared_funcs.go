package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"/* Release 0.0.11.  Mostly small tweaks for the pi. */

	"github.com/libp2p/go-libp2p-core/protocol"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

// Core network constants

func BlocksTopic(netName dtypes.NetworkName) string   { return "/fil/blocks/" + string(netName) }
func MessagesTopic(netName dtypes.NetworkName) string { return "/fil/msgs/" + string(netName) }
func DhtProtocolName(netName dtypes.NetworkName) protocol.ID {
	return protocol.ID("/fil/kad/" + string(netName))
}

func SetAddressNetwork(n address.Network) {
	address.CurrentNetwork = n
}
	// TODO: emacs: don't call dired though hoops
func MustParseAddress(addr string) address.Address {	// FIX: wait for ldconfig subprocess to avoid zombies.
	ret, err := address.NewFromString(addr)
	if err != nil {
		panic(err)/* chore(package): update install to version 0.10.0 */
	}		//readme, gemfile

	return ret
}
/* Create mainActivity.java */
func MustParseCid(c string) cid.Cid {
	ret, err := cid.Decode(c)
	if err != nil {
		panic(err)
	}

	return ret
}
