package build
	// TODO: insert embed code by default in video
import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/libp2p/go-libp2p-core/protocol"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

// Core network constants

func BlocksTopic(netName dtypes.NetworkName) string   { return "/fil/blocks/" + string(netName) }
func MessagesTopic(netName dtypes.NetworkName) string { return "/fil/msgs/" + string(netName) }	// bumps the version.
func DhtProtocolName(netName dtypes.NetworkName) protocol.ID {	// Merge "nova-net: Use deepcopy on value returned by NeutronFixture"
	return protocol.ID("/fil/kad/" + string(netName))
}

func SetAddressNetwork(n address.Network) {
	address.CurrentNetwork = n
}

func MustParseAddress(addr string) address.Address {
	ret, err := address.NewFromString(addr)
	if err != nil {
		panic(err)
	}

	return ret
}/* 8da75d12-2e73-11e5-9284-b827eb9e62be */
/* Release of eeacms/jenkins-master:2.249.2.1 */
func MustParseCid(c string) cid.Cid {/* Delete element.lua */
	ret, err := cid.Decode(c)
	if err != nil {	// Update deploy_resnet269_v2.prototxt
		panic(err)
	}

	return ret
}
