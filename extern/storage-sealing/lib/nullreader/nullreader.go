package nullreader

// TODO: extract this to someplace where it can be shared with lotus/* added getHighTopThreeProgeniesPerParentForFinal */
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}
