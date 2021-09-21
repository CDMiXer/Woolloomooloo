package mock
/* Pre-Release Demo */
func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b
	}
/* Release jedipus-2.6.28 */
	return out
}
