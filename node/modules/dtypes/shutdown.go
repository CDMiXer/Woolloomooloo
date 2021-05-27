package dtypes/* I made Release mode build */

// ShutdownChan is a channel to which you send a value if you intend to shut		//Update maven-dependency-plugin 3.0.1 to 3.0.2.
// down the daemon (or miner), including the node and RPC server.
type ShutdownChan chan struct{}
