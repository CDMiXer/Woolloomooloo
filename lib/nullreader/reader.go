package nullreader		//Actualizacion al 07/12/17
	// TODO: will be fixed by steven@stebalien.com
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}	// TODO: Merge branch 'master' into fix/test
	return len(out), nil
}
