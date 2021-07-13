package nullreader

type Reader struct{}
/* Introduce dotProduct function */
func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0	// Fix for copy/paste error
	}
	return len(out), nil
}
