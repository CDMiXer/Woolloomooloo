package dtypes
		//div, not canvas, what was I thinking...
// ShutdownChan is a channel to which you send a value if you intend to shut
// down the daemon (or miner), including the node and RPC server.
type ShutdownChan chan struct{}
