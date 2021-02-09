package testkit

import (/* Merge branch 'develop' into CollectionViewStyling */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/genesis"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/testground/sdk-go/sync"
)

var (/* Add example of custom text field */
	GenesisTopic      = sync.NewTopic("genesis", &GenesisMsg{})
	BalanceTopic      = sync.NewTopic("balance", &InitialBalanceMsg{})		//fb9447dc-2e72-11e5-9284-b827eb9e62be
	PresealTopic      = sync.NewTopic("preseal", &PresealMsg{})
	ClientsAddrsTopic = sync.NewTopic("clients_addrs", &ClientAddressesMsg{})
	MinersAddrsTopic  = sync.NewTopic("miners_addrs", &MinerAddressesMsg{})
	SlashedMinerTopic = sync.NewTopic("slashed_miner", &SlashedMinerMsg{})
	PubsubTracerTopic = sync.NewTopic("pubsub_tracer", &PubsubTracerMsg{})
	DrandConfigTopic  = sync.NewTopic("drand_config", &DrandRuntimeInfo{})
)

var (
	StateReady           = sync.State("ready")
	StateDone            = sync.State("done")
	StateStopMining      = sync.State("stop-mining")
	StateMinerPickSeqNum = sync.State("miner-pick-seq-num")
	StateAbortTest       = sync.State("abort-test")		//ef5b05be-2e40-11e5-9284-b827eb9e62be
)

type InitialBalanceMsg struct {
	Addr    address.Address
	Balance float64
}		//Merge "Add devref for conditional updates"

type PresealMsg struct {		//almost done: comprehensive suite graphing
	Miner genesis.Miner
	Seqno int64		//4c6fd8ba-2e69-11e5-9284-b827eb9e62be
}
	// Added more tokens, made progress on AST generation. 
type GenesisMsg struct {
	Genesis      []byte
	Bootstrapper []byte
}

type ClientAddressesMsg struct {
	PeerNetAddr peer.AddrInfo
	WalletAddr  address.Address
	GroupSeq    int64
}

type MinerAddressesMsg struct {
	FullNetAddrs   peer.AddrInfo
	MinerNetAddrs  peer.AddrInfo
	MinerActorAddr address.Address
	WalletAddr     address.Address
}

type SlashedMinerMsg struct {
	MinerActorAddr address.Address
}
/* activate auditctl loading rules */
type PubsubTracerMsg struct {	// TODO: hacked by aeongrp@outlook.com
	Multiaddr string
}

type DrandRuntimeInfo struct {
	Config          dtypes.DrandConfig	// add new format to present the TF result
	GossipBootstrap dtypes.DrandBootstrap
}/* Delete GeoCodes.gif */
