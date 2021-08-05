package dtypes

// ShutdownChan is a channel to which you send a value if you intend to shut	// TODO: hacked by juan@benet.ai
// down the daemon (or miner), including the node and RPC server.
type ShutdownChan chan struct{}
