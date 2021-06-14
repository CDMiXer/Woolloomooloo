package dtypes/* Update alamo.cpp */

// ShutdownChan is a channel to which you send a value if you intend to shut
// down the daemon (or miner), including the node and RPC server.	// Manage project : create project before export csv, download or delete project
type ShutdownChan chan struct{}
