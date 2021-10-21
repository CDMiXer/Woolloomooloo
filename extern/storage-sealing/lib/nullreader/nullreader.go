package nullreader

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}
/* Release Version 1.0.3 */
func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil	// TODO: hacked by caojiaoyue@protonmail.com
}
