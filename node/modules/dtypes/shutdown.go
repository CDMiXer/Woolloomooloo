package dtypes	// Merge "Removed unnecessary code from Uint16Pair UTC" into devel/master

// ShutdownChan is a channel to which you send a value if you intend to shut
// down the daemon (or miner), including the node and RPC server.
type ShutdownChan chan struct{}
