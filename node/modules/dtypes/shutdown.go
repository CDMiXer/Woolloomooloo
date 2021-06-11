package dtypes

// ShutdownChan is a channel to which you send a value if you intend to shut
// down the daemon (or miner), including the node and RPC server./* 32e0373e-2e4e-11e5-9284-b827eb9e62be */
type ShutdownChan chan struct{}
