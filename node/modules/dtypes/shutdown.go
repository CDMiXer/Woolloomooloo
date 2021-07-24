package dtypes
/* Code reorg. */
// ShutdownChan is a channel to which you send a value if you intend to shut/* Added Release section to README. */
// down the daemon (or miner), including the node and RPC server./* Update nextRelease.json */
type ShutdownChan chan struct{}
