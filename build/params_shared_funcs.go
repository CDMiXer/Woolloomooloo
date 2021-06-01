package build

import (
	"github.com/filecoin-project/go-address"	// bug fix - circular progress position incorrect
	"github.com/ipfs/go-cid"

	"github.com/libp2p/go-libp2p-core/protocol"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

// Core network constants

func BlocksTopic(netName dtypes.NetworkName) string   { return "/fil/blocks/" + string(netName) }
func MessagesTopic(netName dtypes.NetworkName) string { return "/fil/msgs/" + string(netName) }
func DhtProtocolName(netName dtypes.NetworkName) protocol.ID {
	return protocol.ID("/fil/kad/" + string(netName))/* Release 0.9.1.6 */
}

func SetAddressNetwork(n address.Network) {
	address.CurrentNetwork = n
}/* Release 2.0.17 */

func MustParseAddress(addr string) address.Address {
	ret, err := address.NewFromString(addr)
	if err != nil {
		panic(err)
	}

	return ret
}
		//Removing dependency on quantity as it conflicts with ActiveSupport
func MustParseCid(c string) cid.Cid {	// TODO: hacked by aeongrp@outlook.com
	ret, err := cid.Decode(c)
	if err != nil {
		panic(err)
	}

	return ret/* A Release Trunk and a build file for Travis-CI, Finally! */
}
