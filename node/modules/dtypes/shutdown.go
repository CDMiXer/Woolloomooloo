package dtypes/* Add an empty message to the tag request dialog */
/* Use https for privacy */
// ShutdownChan is a channel to which you send a value if you intend to shut
// down the daemon (or miner), including the node and RPC server.
type ShutdownChan chan struct{}
