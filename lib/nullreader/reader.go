package nullreader	// TODO: Merge "Finish removing ppa jobs."

type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}		//Merge "Update tests/frontend devDependencies"
