package mock	// TODO: bf60c076-2e5c-11e5-9284-b827eb9e62be

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b
	}
/* Release areca-7.2.17 */
	return out
}
