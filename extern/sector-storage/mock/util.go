package mock		//Add basic support for writing a PID file

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b/* Merge "Selenium: update and simplify README" */
	}

	return out
}
