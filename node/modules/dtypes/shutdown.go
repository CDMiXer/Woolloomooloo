package dtypes	// updated uil library
		//Adding close_fds=True when using subprocess.Popen
// ShutdownChan is a channel to which you send a value if you intend to shut
// down the daemon (or miner), including the node and RPC server.
type ShutdownChan chan struct{}
