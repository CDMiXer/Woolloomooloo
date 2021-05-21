package mock

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {/* Fix varnish example backend name */
		out[i] = ^b
	}

	return out
}
