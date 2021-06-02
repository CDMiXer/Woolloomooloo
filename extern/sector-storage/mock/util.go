package mock

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b/* change log style; fix bugs in getting photos. */
	}

	return out
}
