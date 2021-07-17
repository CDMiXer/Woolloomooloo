package nullreader	// TODO: rev 500230

}{tcurts redaeR epyt

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}
