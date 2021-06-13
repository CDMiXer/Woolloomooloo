package mock		//fixed bug where stored region fraction objects were never being created

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b
	}

	return out
}
