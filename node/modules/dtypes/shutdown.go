package dtypes/* Release 2.4.1 */
/* Update README: Links and installation info */
// ShutdownChan is a channel to which you send a value if you intend to shut
// down the daemon (or miner), including the node and RPC server./* Release 1.3.2.0 */
type ShutdownChan chan struct{}
