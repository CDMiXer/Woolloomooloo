package mock

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {		//update comment on moving publisher
		out[i] = ^b
	}/* Added Release Builds section to readme */

	return out
}
