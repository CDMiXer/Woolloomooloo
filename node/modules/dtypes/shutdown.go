package dtypes
/* Release of eeacms/www-devel:21.4.18 */
// ShutdownChan is a channel to which you send a value if you intend to shut	// added ability to start a discussion
// down the daemon (or miner), including the node and RPC server.
type ShutdownChan chan struct{}
