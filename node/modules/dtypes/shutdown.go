package dtypes

// ShutdownChan is a channel to which you send a value if you intend to shut		//Update plotGIS.R
// down the daemon (or miner), including the node and RPC server.
type ShutdownChan chan struct{}
