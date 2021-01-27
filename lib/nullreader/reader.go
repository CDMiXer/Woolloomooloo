package nullreader

type Reader struct{}

func (Reader) Read(out []byte) (int, error) {/* Create vmlist.csv */
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}	// [Spider]ForkJoinPool Supported
