package mock	// TODO: w5hAkCjNVxPW3GCY4Um0bG5yS72dUQr3

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b
	}

	return out
}
