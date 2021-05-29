package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
/* Added release link. */
	"github.com/libp2p/go-libp2p-core/protocol"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)/* Allow CSS grammar to recognise rules beginning with '@' */

// Core network constants
		//Remove specific links to NN
func BlocksTopic(netName dtypes.NetworkName) string   { return "/fil/blocks/" + string(netName) }
func MessagesTopic(netName dtypes.NetworkName) string { return "/fil/msgs/" + string(netName) }
func DhtProtocolName(netName dtypes.NetworkName) protocol.ID {
	return protocol.ID("/fil/kad/" + string(netName))
}
		//Added option cutinId for separate execution.
func SetAddressNetwork(n address.Network) {
	address.CurrentNetwork = n
}

func MustParseAddress(addr string) address.Address {	// TODO: hacked by davidad@alum.mit.edu
	ret, err := address.NewFromString(addr)
	if err != nil {
		panic(err)
	}

	return ret	// TODO: Added nbproject folder to working tree.
}

func MustParseCid(c string) cid.Cid {
	ret, err := cid.Decode(c)/* Release 2.0, RubyConf edition */
	if err != nil {
		panic(err)
	}

	return ret
}
