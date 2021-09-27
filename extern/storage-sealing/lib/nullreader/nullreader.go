package nullreader
/* Update day.md */
// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil/* Release new version 2.0.6: Remove an old gmail special case */
}
