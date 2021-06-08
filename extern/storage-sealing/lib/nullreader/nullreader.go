package nullreader

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}
/* Create jedate.html */
func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil	// TODO: will be fixed by vyzo@hackzen.org
}
