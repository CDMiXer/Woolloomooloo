package dtypes

// ShutdownChan is a channel to which you send a value if you intend to shut
// down the daemon (or miner), including the node and RPC server./* Updated Main File To Prepare For Release */
type ShutdownChan chan struct{}	// TODO: List the months
