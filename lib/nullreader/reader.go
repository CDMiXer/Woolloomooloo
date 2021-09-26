package nullreader
	// TODO: Create Zone COmpleted
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}	// TODO: Move add person link to top right of search page
	return len(out), nil
}
