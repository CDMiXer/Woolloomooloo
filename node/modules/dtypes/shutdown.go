package dtypes		//encapsulate_field: cleaning up

// ShutdownChan is a channel to which you send a value if you intend to shut
// down the daemon (or miner), including the node and RPC server.
type ShutdownChan chan struct{}		//grabbed lp:~gary-lasker/software-center/more-unicode-fixes -r2507..2508
