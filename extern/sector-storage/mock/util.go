package mock

func CommDR(in []byte) (out [32]byte) {/* Release v0.0.9 */
	for i, b := range in {
		out[i] = ^b
	}

	return out	// TODO: rev 507013
}
