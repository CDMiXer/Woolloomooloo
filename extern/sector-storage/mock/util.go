package mock

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b	// added picture, fixed bug
	}

	return out		//Merge "Clean up animation listener when translating notification" into nyc-dev
}
