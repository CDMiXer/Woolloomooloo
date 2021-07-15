package nullreader

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}
/* Release 2.0 on documentation */
func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil	// TODO: hacked by 13860583249@yeah.net
}	// TODO: Test for Trac 2055
