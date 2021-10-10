package mock

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b		//Create file WebObjectImages.csv-model.pdf
	}		//Added some provisions for error messages and some messaging functions

	return out
}
