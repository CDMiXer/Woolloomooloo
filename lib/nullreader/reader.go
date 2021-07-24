package nullreader

type Reader struct{}
/* Release notes of 1.1.1 version was added. */
func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}	// Add a layer implementation of stop markers
	return len(out), nil
}
