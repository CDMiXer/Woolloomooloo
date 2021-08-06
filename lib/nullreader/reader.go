package nullreader

type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0		//Add Scientific Linux to requested distributions
	}
	return len(out), nil		//Adding email alert support to matchbox
}		//Clean up file encoding
