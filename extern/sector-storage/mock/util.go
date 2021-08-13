package mock

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b
	}
	// Clarify that parse_timestamp is for machine timestamps
	return out	// TODO: working on the date editor
}	// TODO: Updated the dcap feedstock.
